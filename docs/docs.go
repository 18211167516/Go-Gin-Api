// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/18211167516/Go-Gin-Api",
        "contact": {
            "name": "baichonghua",
            "email": "18211167516@163.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "用户列表",
                "responses": {
                    "0": {
                        "description": "查询成功",
                        "schema": {
                            "$ref": "#/definitions/models.UserSwagger"
                        }
                    },
                    "20001": {
                        "description": "Token鉴权失败",
                        "schema": {
                            "$ref": "#/definitions/tool.JSONRET"
                        }
                    },
                    "20002": {
                        "description": "Token已超时",
                        "schema": {
                            "$ref": "#/definitions/tool.JSONRET"
                        }
                    },
                    "20004": {
                        "description": "Token错误",
                        "schema": {
                            "$ref": "#/definitions/tool.JSONRET"
                        }
                    },
                    "20005": {
                        "description": "Token参数不能为空",
                        "schema": {
                            "$ref": "#/definitions/tool.JSONRET"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/tool.JSONRET"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.User": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "deleted_at": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "modified_on": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.UserSwagger": {
            "type": "object",
            "properties": {
                "lists": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "tool.JSONRET": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "error_code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:8080",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "go-api 框架",
	Description: "gin-web框架",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
