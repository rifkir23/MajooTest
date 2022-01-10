// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/all": {
            "get": {
                "description": "All Receipt Sea",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "All example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entity.Resi"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/count": {
            "get": {
                "description": "Count Receipt Sea",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "All example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.CountDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CountDTO": {
            "type": "object",
            "properties": {
                "arrivedSoon": {
                    "type": "integer"
                },
                "delay": {
                    "type": "integer"
                },
                "onTheWay": {
                    "type": "integer"
                }
            }
        },
        "entity.Container": {
            "type": "object",
            "properties": {
                "id_container": {
                    "type": "integer"
                },
                "id_rts": {
                    "type": "string"
                },
                "kode": {
                    "type": "string"
                },
                "nomor": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "tanggal_arrived_c": {
                    "type": "string"
                },
                "tanggal_berangkat_c": {
                    "type": "string"
                },
                "tanggal_monitoring_c": {
                    "type": "string"
                },
                "tgl_antri_kapal": {
                    "type": "string"
                },
                "tgl_atur_kapal": {
                    "type": "string"
                },
                "tgl_closing": {
                    "type": "string"
                },
                "tgl_est_dumai": {
                    "type": "string"
                },
                "tgl_eta": {
                    "type": "string"
                },
                "tgl_loading": {
                    "type": "string"
                },
                "tgl_notul": {
                    "type": "string"
                },
                "tgl_pib": {
                    "type": "string"
                }
            }
        },
        "entity.Giw": {
            "type": "object",
            "properties": {
                "barang": {
                    "type": "string"
                },
                "berat": {
                    "type": "string"
                },
                "container": {
                    "$ref": "#/definitions/entity.Container"
                },
                "container_id": {
                    "type": "integer"
                },
                "harga": {
                    "type": "number"
                },
                "harga_jual": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "kurs": {
                    "type": "number"
                },
                "nilai": {
                    "type": "string"
                },
                "nomor": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "remarks": {
                    "type": "string"
                },
                "resi_id": {
                    "type": "integer"
                },
                "supplier": {
                    "type": "integer"
                },
                "tel": {
                    "type": "integer"
                },
                "volume": {
                    "type": "string"
                }
            }
        },
        "entity.Resi": {
            "type": "object",
            "properties": {
                "giw": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Giw"
                    }
                },
                "id_resi": {
                    "type": "integer"
                },
                "konfirmasi_resi": {
                    "type": "string"
                },
                "nomor": {
                    "type": "string"
                },
                "supplier": {
                    "type": "string"
                },
                "tanggal": {
                    "type": "string"
                },
                "tel": {
                    "type": "string"
                }
            }
        },
        "helper.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
