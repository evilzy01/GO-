package defs

import ("")

// Request
type UserCredential struct{
	Username string 'json:"user_name"'
	Pwd string 'json:"pwd"'
}