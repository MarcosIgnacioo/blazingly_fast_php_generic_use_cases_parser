package web_files_manipulation

import (
	"github.com/sunshineplan/node"
	"golang.org/x/net/html"
	"strings"
)

func newTextHtmlNode(text string) *html.Node {
	return (&html.Node{
		Parent:      nil,
		PrevSibling: nil,
		NextSibling: nil,
		Data:        text,
		Type:        html.RawNode, // XD
		Attr:        []html.Attribute{},
	})
}

func removeChildrenWithClassName(tag *node.Node, class string) {
	children := (*tag).Children()
	for i := 0; i < len(children); i++ {
		childClass, exists := children[i].Attrs().Get("class")
		if !exists {
			continue
		}
		if strings.Contains(childClass, class) {
			child := children[i].Raw()
			(*tag).Raw().RemoveChild(child)
		}
	}
}

func removeAllChildren(tag *node.Node) {
	if tag == nil || *tag == nil {
		return
	}
	children := (*tag).Children()
	for i := 0; i < len(children); i++ {
		child := children[i].Raw()
		(*tag).Raw().RemoveChild(child)
	}
}

func insertRawTextBeforeNode(text string, beforeNode node.Node) {
	textNodeToInsert := &html.Node{
		Parent:      nil,
		PrevSibling: nil,
		NextSibling: nil,
		Data:        text,
		Type:        html.RawNode, // XD
		Attr:        []html.Attribute{},
	}
	beforeNode.Raw().InsertBefore(textNodeToInsert, beforeNode.Raw())
}

func setUpAnchors(anchors *[]node.Node, productHref string, productName string) {
	for i := 0; i < len(*anchors); i++ {
		anchor := &(*anchors)[i]
		setAttribute(anchor, "href", productHref)
		setAttribute(anchor, "title", productName)
	}
}

func setUpImages(images *[]node.Node, productImg string, productName string) {
	for i := 0; i < len(*images); i++ {
		img := &(*images)[i]
		setAttribute(img, "src", productImg)
		setAttribute(img, "srcset", productImg)
		setAttribute(img, "alt", productName)
	}
}

func setUpPriceAndName(container *[]node.Node, productPrice string, productName string) {

}

func setAttribute(htmlNode *node.Node, attribute string, value string) {
	attributes := (*htmlNode).Raw().Attr
	// c goat
	for i := 0; i < len(attributes); i++ {
		attr := &attributes[i]
		if attr.Key == attribute {
			attr.Val = value
			break
		}
	}
}

func getData(instruction map[string]string, attribute string) string {
	attributeData := instruction[attribute]
	if attributeData != "" {
		return attributeData
	}
	switch attribute {
	case "href":
		return productHref
	case "productName":
		return productName
	case "img":
		return productThumbnail
	case "price":
		return productPrice
	}
	return attributeData
}
