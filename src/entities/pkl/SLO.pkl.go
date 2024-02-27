// Code generated from Pkl module `srenity.entities.Configuration`. DO NOT EDIT.
package pkl

type SLO struct {
	Name string `pkl:"name"`

	Description string `pkl:"description"`

	Output string `pkl:"output"`

	SLIs []*SLI `pkl:"slis"`
}
