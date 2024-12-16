package web_files_manipulation

import (
	"fmt"
	"os"
	"strings"

	"github.com/sunshineplan/node"
	"golang.org/x/net/html"
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

func removeChildren(tag node.Node) {
	if tag == nil || tag == nil {
		return
	}
	children := (tag).Children()
	for i := 0; i < len(children); i++ {
		child := children[i].Raw()
		(tag).Raw().RemoveChild(child)
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

func removeAllChildrenExceptFirst(tag *node.Node) {
	if tag == nil || *tag == nil {
		return
	}
	children := (*tag).Children()
	for i := 1; i < len(children); i++ {
		child := children[i].Raw()
		(*tag).Raw().RemoveChild(child)
	}
}

func preppendHTMLToNode(text string, beforeNode node.Node) {
	if beforeNode == nil {
		panic("beforeNode nil in preppendHTML")
	}
	if beforeNode.Raw().Parent == nil {
		panic("parent nil in preppendHTML")
	}
	beforeNode.Raw().InsertBefore(newTextHtmlNode(text), beforeNode.Raw())
}

// this doesnt work if the beforenode doesnt have a sibling lol
// tofix
func appendHTMLToNode(text string, where node.Node) {
	if where == nil {
		panic("where nil in appendingHMTL")
	}
	where.Raw().AppendChild(newTextHtmlNode(text))
}

func replaceInnerHTMLFromNode(text string, tag node.Node) {
	if tag == nil {
		panic("tag nil in replacing innerHTML")
	}
	if tag.Raw().Parent == nil {
		panic("parent nil in replacing innerHTML")
	}
	removeChildren(tag)
	tag.Raw().AppendChild(newTextHtmlNode(text))
}

func replaceOuterHTMLFromNode(outerHTML string, tag node.Node) {
	if tag == nil {
		panic("tag nil in replacing outerHTML")
	}
	if tag.Raw().Parent == nil {
		panic("parent nil in replacing outerHTML")
	}
	tag.Raw().InsertBefore(newTextHtmlNode(outerHTML), tag.Raw())
	tag.Parent().Raw().RemoveChild(tag.Raw())
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
			return
		}
	}
	newAttribute := html.Attribute{Key: attribute, Val: value}
	(*htmlNode).Raw().Attr = append(attributes, newAttribute)
}

func appendAttribute(htmlNode *node.Node, attribute string, value string) {
	attributes := (*htmlNode).Raw().Attr
	// c goat
	for i := 0; i < len(attributes); i++ {
		attr := &attributes[i]
		if attr.Key == attribute {
			attr.Val += " " + value
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

func dirToJson() map[string][]map[string]string {
	return map[string][]map[string]string{
		"DIRNAME": {
			{
				"class":        "html",
				"htmlToInsert": merchHeader,
			},
			{
				"class":      "product_item_merch",
				"foreach":    merchForeachWrapper,
				"classPrice": "product_price",
				"className":  "product_name",
			},
		},
	}
}

func setUpTags(tags *[]node.Node, tagAttribute Attribute) {
	for i := 0; i < len(*tags); i++ {
		tag := &(*tags)[i]
		setAttribute(tag, tagAttribute.Name, tagAttribute.Value)
	}
}

func HTMLManipulation(instruction Instruction, targetDiv node.Node) {
	if instruction.InnerHTML != "" {
		replaceInnerHTMLFromNode(instruction.InnerHTML, targetDiv)
	}

	if instruction.PrependHTML != "" {
		preppendHTMLToNode(instruction.PrependHTML, targetDiv)
	}

	if instruction.AppendHTML != "" {
		appendHTMLToNode(instruction.AppendHTML, targetDiv)
	}

	if instruction.OuterHTML != "" {
		replaceOuterHTMLFromNode(instruction.OuterHTML, targetDiv)
	}
}

func attributesManipulation(instruction Instruction, classToSearch string, targetDiv node.Node) {
	for _, tagAttribute := range instruction.TagsAttributes {
		var tags []node.Node
		if tagAttribute.Tag != "" {
			tags = QuerySelectorAll(targetDiv, tagAttribute.Tag)
		} else {
			tags = []node.Node{targetDiv}
		}
		if tags == nil || len(tags) == 0 {
			panic(fmt.Sprintf("%s are nil or zero len in this instruction %s", classToSearch, tagAttribute.Tag))
		}
		setUpTags(&tags, tagAttribute.Attr)
	}
}

func innerHTMLManipulation(instruction Instruction, classToSearch string, targetDiv node.Node) {
	for _, htmlReplacement := range instruction.InnerHtmlReplacements {
		oldHtmlContainer := targetDiv.Find(node.Descendant, nil, node.Class(htmlReplacement.ClassName))
		if oldHtmlContainer == nil {
			panic(fmt.Sprintf(" oldHtmlContainer is nil\nthis is the classname we are looking for `%s`", htmlReplacement.ClassName))
		}
		removeAllChildren(&oldHtmlContainer)
		replaceInnerHTMLFromNode(htmlReplacement.HTML, oldHtmlContainer)
	}
}

func constructHTML(parent *node.Node, targetDiv node.Node, instruction Instruction, shouldRemoveChildren bool) {
	// solamente quiero vacaciones
	if !instruction.IsParent {
		*parent = targetDiv.Parent()
	} else {
		*parent = targetDiv
	}
	htmlToInsert := targetDiv.HTML()
	if instruction.ForEach != "" {
		htmlToInsert = fmt.Sprintf(instruction.ForEach, htmlToInsert)
	}
	if shouldRemoveChildren {
		removeAllChildren(parent)
		// replaceInnerHTMLFromNode(htmlToInsert, targetDiv)
	}
	(*parent).Raw().AppendChild(newTextHtmlNode(htmlToInsert))
}

func buildPHPFile(file *File, doc node.Node) {
	phpFileName := strings.Replace(file.filePath, "html", "php", -1)
	os.Rename(file.filePath, phpFileName)
	file.filePath = phpFileName
	os.WriteFile(phpFileName, prepareHTMLForFile(doc), 0777)
	fmt.Println(file.filePath)
}

func getBasicInstruction(className string, forEachWrapper string, priceClassName string, productNameClassName string, productPriceHTML string, productNameHTML string, imgs ...TagAttribute) Instruction {
	imgSrc := productImgSrcTagAtrr
	imgSrcSet := productImgSrcSetTagAtrr
	if len(imgs) > 1 {
		imgSrc = imgs[0]
		imgSrcSet = imgs[1]
	}
	return Instruction{
		Class:                   className,
		ShouldRemoveAllChildren: true,
		ForEach:                 forEachWrapper,
		TagsAttributes: []TagAttribute{
			productAnchorHrefTagAtrr,
			productAnchorTitleTagAtrr,
			productImgAltTagAtrr,
			imgSrc,
			imgSrcSet,
		},
		InnerHtmlReplacements: []HTMLReplacement{
			HTMLReplacement{
				ClassName: priceClassName,
				HTML:      productPriceHTML,
			},
			HTMLReplacement{
				ClassName: productNameClassName,
				HTML:      productNameHTML,
			},
		},
	}
}

func insertCartMobile(doc node.Node, nestedLevel int) {
	cartMobile := fmt.Sprintf("<?php include \"%slayouts/cart_mobile.template.php\"; ?>\n", getNestedPath(nestedLevel))
	separatorCart := doc.Find(node.Descendant, nil, node.Class("separator_initial_cart"))
	preppendHTMLToNode(cartMobile, separatorCart)
}

func getNestedPath(nestedLevel int) string {
	nesting := ""
	if nestedLevel > 1 {
		for i := 0; i < nestedLevel; i++ {
			nesting += "../"
		}
	}
	return nesting
}

func QuerySelectorAll(from node.Node, fullQuery string) []node.Node {
	queries := strings.Split(fullQuery, " ")
	return qRecurse(from, queries, 0)
}

func qRecurse(from node.Node, queries []string, queryIdx int) []node.Node {
	currentQuery := queries[queryIdx]
	// pre
	fetchedNodes := searchElements(currentQuery, from)
	if queryIdx == len(queries)-1 {
		return fetchedNodes
	}
	// recurse
	results := make([]node.Node, 1)
	for _, fetchedNode := range fetchedNodes {
		results = append(results, qRecurse(fetchedNode, queries, queryIdx+1)...)
	}
	return results
}

func QuerySelector(document node.Node, query string) node.Node {
	queries := parseQuery(query)
	if len(queries) <= 0 {
		return nil
	}
	elements := searchElements(queries[0], document)

	if len(elements) == 0 {
		panic(fmt.Sprintf(" elements is nil when trying to select\nthis query: `%s`\nprobably forgot to put a `.` for the className", query))
		// return nil
	}
	needleElement := elements[0]
	for i := 1; i < len(queries); i++ {
		for _, element := range elements {
			needleElement = searchElement(queries[i], element)
			if needleElement != nil {
				break
			}
		}
	}
	return needleElement
}

func parseQuery(query string) []string {
	return strings.Split(query, " ")
}

func getClassNameOrTagName(query string) string {
	if query[0] == '.' {
		return query[1:]
	} else {
		return query
	}
}

func searchElements(name string, where node.Node) []node.Node {
	if name[0] == '.' {
		return where.FindAll(node.Descendant, nil, node.Class(name[1:]))
	} else {
		return where.FindAll(node.Descendant, node.Tag(name))
	}
}

func searchElement(name string, where node.Node) node.Node {
	if name[0] == '.' {
		return where.Find(node.Descendant, nil, node.Class(name[1:]))
	} else {
		return where.Find(node.Descendant, node.Tag(name))
	}
}

// func newInstruction(directory string, productClass string, header string, forEachWrapper string, classNames ...string) map[string][]Instruction {
// 	productPriceClass := "productPrice"
// 	productNameClass := "productName"
// 	if len(classNames) > 1 {
// 		productPriceClass = classNames[0]
// 		productNameClass = classNames[1]
// 	}
// 	return map[string][]Instruction

// }
