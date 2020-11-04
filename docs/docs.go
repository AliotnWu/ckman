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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/ck/table": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改表",
                "summary": "修改表",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AlterCkTableReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":5003,\"msg\":\"更改ClickHouse表失败\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建表",
                "summary": "创建表",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateCkTableReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":5001,\"msg\":\"创建ClickHouse表失败\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除表",
                "summary": "删除表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "table name",
                        "name": "tableName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":5002,\"msg\":\"删除ClickHouse表失败\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/deploy/ck": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "部署clickhouse",
                "summary": "部署clickhouse",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DeployCkReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":5016,\"msg\":\"检查组件启动状态失败\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "登陆",
                "summary": "登陆",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":5032,\"msg\":\"用户密码验证失败\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/package": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取安装包列表",
                "summary": "获取安装包列表",
                "responses": {
                    "200": {
                        "description": "{\"code\":5005,\"msg\":\"获取安装包列表失败\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "上传安装包",
                "consumes": [
                    "multipart/form-data"
                ],
                "summary": "上传安装包",
                "parameters": [
                    {
                        "type": "file",
                        "description": "安装包",
                        "name": "package",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":5004,\"msg\":\"上传安装包失败\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除包",
                "summary": "删除包",
                "parameters": [
                    {
                        "type": "string",
                        "description": "package name",
                        "name": "packageName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":5002,\"msg\":\"删除ClickHouse表失败\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AlterCkTableReq": {
            "type": "object",
            "properties": {
                "add": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CkTableNameTypeAfter"
                    }
                },
                "drop": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "modify": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CkTableNameType"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.CkCluster": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "shards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CkShard"
                    }
                }
            }
        },
        "model.CkDeployConfig": {
            "type": "object",
            "properties": {
                "clusters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CkCluster"
                    }
                },
                "password": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                },
                "zkServers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ZkNode"
                    }
                }
            }
        },
        "model.CkReplica": {
            "type": "object",
            "properties": {
                "host": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        },
        "model.CkShard": {
            "type": "object",
            "properties": {
                "replicas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CkReplica"
                    }
                }
            }
        },
        "model.CkTableNameType": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.CkTableNameTypeAfter": {
            "type": "object",
            "properties": {
                "after": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.CkTablePartition": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "policy": {
                    "type": "integer"
                }
            }
        },
        "model.CreateCkTableReq": {
            "type": "object",
            "properties": {
                "fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CkTableNameType"
                    }
                },
                "name": {
                    "type": "string"
                },
                "order": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "partition": {
                    "$ref": "#/definitions/model.CkTablePartition"
                }
            }
        },
        "model.DeployCkReq": {
            "type": "object",
            "properties": {
                "clickHouse": {
                    "$ref": "#/definitions/model.CkDeployConfig"
                },
                "directory": {
                    "type": "string"
                },
                "hosts": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "packages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "password": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "model.LoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.ZkNode": {
            "type": "object",
            "properties": {
                "host": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "token",
            "in": "header"
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
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "",
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
