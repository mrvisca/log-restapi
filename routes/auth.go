package routes

import (
	"fmt"
	"log-restapi/config"
	"log-restapi/models"
	"net/http"
	"os"
	"time"

	"github.com/danilopolani/gocialite/structs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

func CheckToken(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Sukses login, Selamat Datang",
		"status":  "Login Sukses!",
	})
}

// Login godoc
// @Summary	Login akun
// @Description	Redirect ke halaman github.com / google.com untuk autentitakasi dengan akun github / google
// @Tags autentikasi
// @Param provider path string true "Provider (google / github)"
// @Produce	application/json
// @Success	200
// @Router /auth/github [get]
func RedirectHandler(c *gin.Context) {
	provider := c.Param("provider")

	// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
	// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     os.Getenv("CLIENT_ID_GH"),
			"clientSecret": os.Getenv("CLIENT_SECRET_GH"),
			"redirectURL":  os.Getenv("AUTH_REDIRECT_URL") + "/github/callback",
		},
		"google": {
			"clientID":     os.Getenv("CLIENT_ID_G"),
			"clientSecret": os.Getenv("CLIENT_SECRET_G"),
			"redirectURL":  os.Getenv("AUTH_REDIRECT_URL") + "/google/callback",
		},
	}

	providerScopes := map[string][]string{
		"github": []string{},
		"google": []string{},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := config.Gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}

func CallbackHandler(c *gin.Context) {
	// Ambil paramter kueri untuk status dan kode
	state := c.Query("state")
	code := c.Query("code")
	provider := c.Param("provider")

	// Handle callback and check for errors
	user, token, err := config.Gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	var newUser = getOrRegisterUser(provider, user)
	var newToken = createToken(&newUser)

	c.JSON(200, gin.H{
		"data":     newUser,
		"token":    newToken,
		"gh_token": token,
		"message":  "Berhasil Login Aplikasi",
	})
}

func getOrRegisterUser(provider string, user *structs.User) models.User {
	var userData models.User

	config.DB.Where("provider = ? AND social_id = ?", provider, user.ID).First(&userData)

	sekarang := time.Now()

	if userData.ID == 0 {
		newUser := models.User{
			Username: user.Username,
			Fullname: user.FullName,
			Email:    user.Email,
			SocialId: user.ID,
			Provider: provider,
			Avatar:   user.Avatar,
			Role:     true,
			Limit:    1,
			IsMimin:  false,
		}
		config.DB.Create(&newUser)

		newOffice := models.Office{
			UserId: newUser.ID,
			Name:   "Kantor Milik " + newUser.Username,
			Email:  user.Email,
			Alamat: "Jakarta, Indonesia",
			Telpon: "082140466335",
			Join:   sekarang.Format("2006-01-02"),
		}
		config.DB.Create(&newOffice)

		return newUser
	} else {
		return userData
	}
}

func ProfileUser(c *gin.Context) {
	user_id := int(c.MustGet("jwt_user_id").(float64))
	var user models.User

	// Eager loading (mengakses data pada beberapa tabel yang memiliki relasi untuk ditampilakan bersamaan tanpa harus query 1 per satu)
	item := config.DB.Where("id = ?", user_id).Preload("Office", "user_id = ?", user_id).Find(&user) // Office dari user struct Offices untuk penghubung

	c.JSON(200, gin.H{
		"status": "Sukses mengakses halaman profil",
		"data":   item,
	})
}

func createToken(user *models.User) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"user_role": user.Role,
		"limit":     user.Limit,
		"is_mimin":  user.IsMimin,
		"exp":       time.Now().AddDate(0, 0, 7).Unix(),
		"iat":       time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
	}

	return tokenString

}
