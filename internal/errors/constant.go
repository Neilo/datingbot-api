package errors

import "net/http"

//BadRequest errors returns when user make a be reg
var BadRequest = Error{MSG: "Bad Request", Status: http.StatusBadRequest}

//BadMethod errors when method not implemented
var BadMethod = Error{MSG: "Bad Method", Status: http.StatusNotFound}

//DataBaseError errors when database error
var DataBaseError = Error{MSG: "Internal DataBase error", Status: http.StatusInternalServerError}
