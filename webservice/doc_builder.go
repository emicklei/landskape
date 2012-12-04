package webservice

import (

)
type DocBuilder struct {
	
}
func (self DocBuilder) Comment(comment string) {}
func (self DocBuilder) QueryParam(name, comment string) {}
func (self DocBuilder) PathParam(name, comment string) {}
func (self DocBuilder) Consumes(mimetypes ...string) {}
func (self DocBuilder) Produces(mimetypes ...string) {}
