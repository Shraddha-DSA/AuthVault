package storage

import "authvault/models"

var Users = map[string]models.User{}
var Blacklist = map[string]bool{}
