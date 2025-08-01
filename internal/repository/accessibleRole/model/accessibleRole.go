package model

import (

)


type EndpointRole struct {
	Endpoint string  `db:"endpoint"`
	Role string `db:role`
}