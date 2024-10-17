package parser

import "piblog/src/model"

type IParser interface {
	ParseDocument(input string) model.Document
}
