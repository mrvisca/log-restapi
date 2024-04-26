// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support (Mr.Visca)",
            "url": "https://resume.mrvisca.tech",
            "email": "bimaputra@mrvisca.tech"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/github": {
            "get": {
                "description": "Redirect ke halaman github.com / google.com untuk autentitakasi dengan akun github / google",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "autentikasi"
                ],
                "summary": "Login akun",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider (google / github)",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1/",
	Schemes:          []string{},
	Title:            "Service API Documentation",
	Description:      "Dokumentasi Layanan API untuk monitor log aktivitas",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
