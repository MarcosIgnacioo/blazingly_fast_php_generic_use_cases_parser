package web_files_manipulation

import (
	"fmt"
	"os"
	// "os/exec"
	"strings"

	"github.com/sunshineplan/node"
	"golang.org/x/net/html"
)

func p(a ...any) {
	fmt.Println(a...)
}

func f(fString string, a ...any) string {
	return fmt.Sprintf(fString, a...)
}

func fp(fString string, a ...any) {
	fmt.Printf(fString, a...)
}

func panik(fString string, a ...any) {
	panic(f(fString, a...))
}

func NewTextHtmlNode(text string) *html.Node {
	return (&html.Node{
		Parent:      nil,
		PrevSibling: nil,
		NextSibling: nil,
		Data:        text,
		Type:        html.RawNode, // XD
		Attr:        []html.Attribute{},
	})
}

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
		classes := strings.Split(childClass, " ")
		hasClass := false
		for _, innerClass := range classes {
			if innerClass == class {
				hasClass = true
				break
			}
		}
		if hasClass {
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

func RemoveAllChildrenExceptThis(parent node.Node, exception node.Node) {
	if parent == nil {
		return
	}
	children := parent.Children()
	rawException := exception.Raw()
	for i := 0; i < len(children); i++ {
		child := children[i].Raw()
		// the 'better' thing to do would be to make `HandleHTMLModifications`
		// to return each html.node corresponding to all the `newTextHtmlNode`
		// and pass them here so we dont delete them too
		// butttttt im too lazy im so sorry
		if child == rawException || strings.Contains(children[i].HTML(), "<?php") {
			continue
		}
		parent.Raw().RemoveChild(child)
	}
}

func removeAllChildrenExceptFirstWithClassName(tag *node.Node, className string) {
	if tag == nil || *tag == nil {
		return
	}
	children := (*tag).Children()
	firstLock := false
	for i := 0; i < len(children); i++ {
		child := children[i].Raw()
		childClassName, _ := children[i].Attrs().Get("class")
		containsClass := strings.Contains(childClassName, className)
		if containsClass && firstLock {
			(*tag).Raw().RemoveChild(child)
		} else if containsClass {
			firstLock = true
		}
	}
}

// una alternativa a esto seria hacer una copia del primero y borrar a todos los hijos del contenedor padre, luego reemplazarlo con el nodo alterado pero  no seee
// Remove other html elements with the same className except the first one of the parent container. Completly unsafe and relies on the fact the user provides always the first one as the one to use
func RemoveElementsWithClassNameExceptFirst(pivotNode *node.Node, query string, shouldUsePivotParent bool) {
	if pivotNode == nil || *pivotNode == nil {
		panik("pivotNode removing node is nil")
	}

	var parent node.Node
	if shouldUsePivotParent {
		parent = (*pivotNode).Parent()
	} else {
		parent = *pivotNode
	}

	if parent == nil {
		panik("parent of `removing class`: %s is nil", query)
	}

	children := QuerySelectorAll(parent, query)

	if children == nil {
		panik("children of `removing class`: %s are nil", query)
	}

	for i := 1; i < len(children); i++ {
		child := children[i].Raw()
		childClassName, exists := children[i].Attrs().Get("class")
		if !exists {
			continue
		}
		containsClass := HasSameClass(childClassName, query)
		if containsClass {
			parent.Raw().RemoveChild(child)
		}
	}
}

func DeleteOtherElementCopies(targets []node.Node) {
	if len(targets) < 1 {
		return
	}
	parent := targets[0].Parent()
	for i := 1; i < len(targets); i++ {
		child := targets[i].Raw()
		if child.Parent != parent.Raw() {
			fakeChildClass, _ := targets[i].Attrs().Get("class")
			expectedToBeParentClass, _ := parent.Attrs().Get("class")
			// realParentClass, _ := targets[i].Parent().Attrs().Get("class")
			p("\n************************************************************************************************************************************************")
			// expectedParentId, _ := targets[0].Attrs().Get("id")
			// realParentId, _ := targets[i].Parent().Attrs().Get("id")
			// p("expected parent id:", expectedParentId)
			// p("real parent id:", realParentId)
			fp("\nthere a nested elements with the same classname WITHOUT THE SAME PARENT!!! \n\n\"fake child\":\n\t`%s`\n\n\"expected to be parent\":\n\t`%s`\n", fakeChildClass, expectedToBeParentClass)
			p("************************************************************************************************************************************************\n")
			// fp("there a nested elements with the same classname WITHOUT THE SAME PARENT!!! \n\"child\":\n%s\n\"parent\":\n%s", targets[i].HTML(), parent.HTML())
			continue
		}
		parent.Raw().RemoveChild(child)
	}
}

// thisClass is the class of the element we are analizing
//
// thisClass = `ed-element ed-container Foto-item product_item_coffe`
//
// thatClass is the class we are looking for
//
// thatClass = `ed-element`
func HasSameClass(thisClass string, thatClass string) bool {
	if thisClass == "" {
		return false
	}
	thatClass = parseClass(thatClass)
	allClasses := strings.Split(thisClass, " ")
	for _, class := range allClasses {
		if parseClass(class) == thatClass {
			return true
		}
	}

	return false
}

func parseClass(class string) string {
	if len(class) == 0 {
		return ""
	}
	if class[0] == '.' {
		return class[1:]
	}
	return class
}

func newHTMLReplacement(className string, html string) HTMLReplacement {
	return HTMLReplacement{
		ClassName: className,
		HTML:      html,
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
			return
		}
	}
	newAttribute := html.Attribute{Key: attribute, Val: value}
	(*htmlNode).Raw().Attr = append(attributes, newAttribute)
}

func setUpTags(tags *[]node.Node, tagAttribute Attribute, shouldAppendAttributes bool) {
	attrChange := setAttribute
	if shouldAppendAttributes {
		attrChange = appendAttribute
	}
	for i := 0; i < len(*tags); i++ {
		tag := &(*tags)[i]
		attrChange(tag, tagAttribute.Name, tagAttribute.Value)
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
		setUpTags(&tags, tagAttribute.Attr, instruction.ShouldAppendAttributes)
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

func constructHTML(parent *node.Node, targetDiv node.Node, instruction Instruction, className string) {
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
	if instruction.ShouldRemoveAllChildren {
		removeAllChildren(parent)
		// replaceInnerHTMLFromNode(htmlToInsert, targetDiv)
	}
	if instruction.ShouldRemoveAllChildrenExceptFirst {
		// removeAllChildrenExceptFirst(parent)
		removeAllChildrenExceptFirstWithClassName(parent, className)
	}

	if instruction.ShouldRemoveTagsWithSameClassName {
		removeChildrenWithClassName(parent, className)
	}

	if instruction.ShouldReplaceOld {
		(*parent).Raw().InsertBefore(newTextHtmlNode(htmlToInsert), targetDiv.Raw())
		(*parent).Raw().RemoveChild(targetDiv.Raw())
	} else {
		(*parent).Raw().AppendChild(newTextHtmlNode(htmlToInsert))
	}
}

func buildPHPFile(file *File, doc node.Node) {
	phpFileName := strings.Replace(file.filePath, "html", "php", -1)
	os.Rename(file.filePath, phpFileName)
	file.filePath = phpFileName
	os.WriteFile(phpFileName, prepareHTMLForFile(doc), 0777)
	// pretty-php test.php
	// phpFormat := exec.Command("pretty-ph.phar", phpFileName)
	// err := phpFormat.Start()
	// if err != nil {
	// 	fp("error php formating file `%s`", phpFileName)
	// }
	// // prettier --parser html
	// htmlFormat := exec.Command("prettier", f("--write --parser html %s"), phpFileName)
	// err = htmlFormat.Start()
	// if err != nil {
	// 	fp("error html formating file `%s`", phpFileName)
	// }
	fmt.Println(file.filePath)
	// p(IDS)
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

func insertScripts(doc node.Node) {
	body := QuerySelector(doc, "body")
	head := QuerySelector(doc, "head")

	InsertBeforeLastChild(scriptsBody, &body)
	InsertBeforeLastChild(scriptsHead, &head)
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
	results := qRecurse(from, queries, 0)
	numberOfNotNilElements := 0
	for _, result := range results {
		if result != nil {
			numberOfNotNilElements++
		}
	}
	notNilResults := make([]node.Node, numberOfNotNilElements)
	i := 0
	for _, result := range results {
		if result != nil {
			notNilResults[i] = result
			i++
		}
	}
	return notNilResults
}

func qRecurse(from node.Node, queries []string, queryIdx int) []node.Node {
	currentQuery := queries[queryIdx]
	// pre
	// fetchedNodes := searchElements(currentQuery, from)
	var fetchedNodes []node.Node
	if strings.Contains(currentQuery, "=") {
		tag, attribute := ParseAttributeQuery(currentQuery)
		fetchedNodes = SearchAllByAttribute(tag, from, attribute)
	} else {
		fetchedNodes = searchElements(currentQuery, from)
	}
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

	var elements []node.Node
	if strings.Contains(queries[0], "=") {
		tag, attribute := ParseAttributeQuery(queries[0])
		elements = SearchAllByAttribute(tag, document, attribute)
	} else {
		elements = searchElements(queries[0], document)
	}

	if len(elements) == 0 {
		p(document.HTML())
		panic(fmt.Sprintf(" elements is nil when trying to select\nthis query: `%s`\nprobably forgot to put a `.` for the className", query))
		// return nil
	}
	needleElement := elements[0]
	for i := 1; i < len(queries); i++ {
		for _, element := range elements {
			if strings.Contains(queries[i], "=") {
				tag, attribute := ParseAttributeQuery(queries[i])
				needleElement = SearchByAttribute(tag, element, attribute)
			} else {
				needleElement = searchElement(queries[i], element)
			}
			// needleElement = searchElement(queries[i], element)
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

func searchElements(name string, where node.Node, params ...any) []node.Node {
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

func generateInputTagAttributes(fields map[string][]string, tag string, header string) []Instruction {
	attrsLen := len(fields)
	instructions := make([]Instruction, attrsLen+1)
	index := 0
	instructions[index] = Instruction{
		Class:       "html",
		PrependHTML: header,
	}
	index++
	for k, v := range fields {
		instructions[index] =
			Instruction{
				Class: k,
				TagsAttributes: []TagAttribute{
					TagAttribute{
						Tag: tag,
						Attr: Attribute{
							Name:  v[0],
							Value: v[1],
						},
					},
				},
			}
		index++
	}
	return instructions
}

func CreateAttribute(name string, value string) Attribute {
	return Attribute{
		Name:  name,
		Value: value,
	}
}

func CreateTagAttribute(name string, value string, params ...any) TagAttribute {
	tagName := ""
	if len(params) > 0 {
		tagName = params[0].(string)
	}
	return TagAttribute{
		// with empty tag we use the Target div from above
		Tag:  tagName,
		Attr: CreateAttribute(name, value),
	}
}

func newHTMLReplacements(htmlReplacements map[string]string) []HTMLReplacement {
	htmlReplacementsLen := len(htmlReplacements)
	replacements := make([]HTMLReplacement, htmlReplacementsLen)
	index := 0
	for k, v := range htmlReplacements {
		replacements[index] = newHTMLReplacement(k, v)
		index++
	}
	return replacements
}

// example of well formed arguments
// `checkout_phone`,
// `input`,
// `name="phone"`,
func CreateInstructionReplacementForAttributeOfTag(className string, attribute string, tagName string, params ...any) Instruction {
	if len(attribute) < 1 {
		panik("attribute empty `%s`", attribute)
	}
	attrData := strings.Split(attribute, "=")
	if className[0] != '.' {
		className = f(".%s", className)
	}

	if len(attrData) < 2 {
		panik(" AttrData is malformed `%s`\nmaybe missing `=`", attribute)
	}

	shouldRemoveChildren := false
	if len(params) > 1 {
		shouldRemoveChildren = true
	}

	attributeName := attrData[0]
	attributeValue := strings.ReplaceAll(attrData[1], "\"", "")

	return Instruction{
		Class:                   className,
		ShouldRemoveAllChildren: shouldRemoveChildren,
		TagsAttributes: []TagAttribute{
			TagAttribute{
				// with empty tag we use the Target div from above
				Tag:  tagName,
				Attr: CreateAttribute(attributeName, attributeValue),
			},
		},
	}
}

// new API

func SearchByAttribute(query string, doc node.Node, attribute string) node.Node {
	name, value := ParseAttribute(attribute)
	matches := QuerySelectorAll(doc, query)
	for _, match := range matches {
		attr, hasIt := match.Attrs().Get(name)
		if !hasIt {
			continue
		}
		if attr == value {
			return match
		}
	}
	return nil
}

func SearchAllByAttribute(query string, doc node.Node, attribute string) []node.Node {
	name, value := ParseAttribute(attribute)
	matches := QuerySelectorAll(doc, query)
	founds := make([]node.Node, 12)
	total := 0
	for _, match := range matches {
		attr, hasIt := match.Attrs().Get(name)
		if !hasIt {
			continue
		}
		if attr == value {
			founds = append(founds, match)
			total++
		}
	}
	notNilFounds := make([]node.Node, total)
	index := 0
	for _, found := range founds {
		if found != nil {
			notNilFounds[index] = found
			index++
		}
	}
	return notNilFounds
}

func PrependHTMLToNode(text string, where node.Node) {
	if where == nil {
		panic("where nil in preppendHTML")
	}
	InsertBefore(newTextHtmlNode(text), where.Raw())
}

func AppendHTMLToNode(text string, where node.Node) {
	if where == nil {
		panic("where nil in appendingHMTL")
	}
	InsertAfter(newTextHtmlNode(text), where.Raw())
}

func ReplaceInnerHTMLFromNode(text string, tag node.Node) {
	if tag == nil {
		panic("tag nil in replacing innerHTML")
	}
	if tag.Raw().Parent == nil {
		panic("parent nil in replacing innerHTML")
	}
	removeChildren(tag)
	tag.Raw().AppendChild(newTextHtmlNode(text))
}

func ReplaceOuterHTMLFromNode(outerHTML string, tag node.Node) {
	if tag == nil {
		panic("tag nil in replacing outerHTML")
	}
	if tag.Raw().Parent == nil {
		panic("parent nil in replacing outerHTML")
	}
	// tag.Parent().Parent().Parent().Parent().Raw().InsertBefore(newTextHtmlNode(outerHTML), tag.Raw())
	// tag.Raw().InsertBefore(newTextHtmlNode(outerHTML), tag.Raw())
	InsertBefore(newTextHtmlNode(outerHTML), tag.Raw())
	tag.Parent().Raw().RemoveChild(tag.Raw())
}

func HandleHTMLModifications(modification Modification, target node.Node) {
	if modification.InnerHTML != "" {
		ReplaceInnerHTMLFromNode(modification.InnerHTML, target)
	}

	if modification.PrependHTML != "" {
		PrependHTMLToNode(modification.PrependHTML, target)
	}

	if modification.AppendHTML != "" {
		AppendHTMLToNode(modification.AppendHTML, target)
	}

	if modification.OuterHTML != "" {
		ReplaceOuterHTMLFromNode(modification.OuterHTML, target)
	}
}

func InsertBefore(newChild, oldChild *html.Node) {
	if newChild.Parent != nil || newChild.PrevSibling != nil || newChild.NextSibling != nil {
		panic("html: InsertBefore called for an attached child Node")
	}
	var prev, next *html.Node
	n := oldChild.Parent

	if n == nil {
		panik("html: oldChild parent is nil")
	}
	if oldChild != nil {
		prev, next = oldChild.PrevSibling, oldChild
	} else {
		prev = n.LastChild
	}
	if prev != nil {
		prev.NextSibling = newChild
	} else {
		n.FirstChild = newChild
	}
	if next != nil {
		next.PrevSibling = newChild
	} else {
		n.LastChild = newChild
	}
	newChild.Parent = n
	newChild.PrevSibling = prev
	newChild.NextSibling = next
}

func InsertAfter(newChild, oldChild *html.Node) {
	if (*newChild).Parent != nil || (*newChild).PrevSibling != nil || (*newChild).NextSibling != nil {
		panic("html: InsertAfter called for an attached child Node")
	}

	n := oldChild.Parent
	if n == nil {
		panik("html: oldChild parent is nil")
	}

	var prev, next *html.Node
	if oldChild != nil {
		prev, next = oldChild, oldChild.NextSibling
	} else {
		prev = n.LastChild
	}
	if prev != nil {
		prev.NextSibling = newChild
	} else {
		n.FirstChild = newChild
	}
	if next != nil {
		next.PrevSibling = newChild
	} else {
		n.LastChild = newChild
	}
	newChild.Parent = n
	newChild.PrevSibling = prev
	newChild.NextSibling = next
}

func InsertBeforeLastChild(rawText string, parent *node.Node) {
	if parent == nil {
		panik(" parent is nil  \n")
	}
	lastChild := (*parent).LastChild()
	if lastChild == nil {
		panik(" last child is nil in parent \n%s", (*parent).HTML())
	}
	InsertAfter(NewTextHtmlNode(rawText), lastChild.Raw())
}

func ParseAttributeQuery(attributeQuery string) (tagName string, attribute string) {
	attrData := strings.Split(attributeQuery, "[")
	tagName = attrData[0]
	attribute = strings.ReplaceAll(attrData[1], "]", "")
	return
}

func ParseAttribute(attribute string) (attributeName string, attributeValue string) {
	attrData := strings.Split(attribute, "=")
	attributeName = attrData[0]
	attributeValue = strings.ReplaceAll(attrData[1], "\"", "")
	return
}

// POSIBLE REFACTOR
// change targetDiv for whole document because if i wanna change multiple
// tags attributes that arent on the same container
// the ones outside the container where the first one came well wont have the changes
// this would be like a really weird edge but still possible
func AttributesChanges(modification Modification, targetDiv node.Node) {
	for _, tagAttribute := range modification.AttributesChanges {
		var tags []node.Node
		if tagAttribute.Query != "" {
			tags = QuerySelectorAll(targetDiv, tagAttribute.Query)
		} else {
			tags = []node.Node{targetDiv}
		}
		if tags == nil || len(tags) == 0 {
			panic(fmt.Sprintf("%s are nil or zero len in this tagAttribute `%s`", tagAttribute.Query))
		}
		StoreID(tags[0], tagAttribute.Query)
		SetUpTags(&tags, tagAttribute.Attribute, tagAttribute.Mode)
	}
}

func SetUpTags(tags *[]node.Node, tagAttribute Attribute, mode AttributeMode) {
	attrChange := setAttribute
	if mode == APPEND_ATTRIBUTE {
		attrChange = appendAttribute
	}
	for i := 0; i < len(*tags); i++ {
		tag := &(*tags)[i]
		attrChange(tag, tagAttribute.Name, tagAttribute.Value)
	}
}

func SetAttribute(htmlNode *node.Node, attribute string, value string) {
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

func AppendAttribute(htmlNode *node.Node, attribute string, value string) {
	attributes := (*htmlNode).Raw().Attr
	// c goat
	for i := 0; i < len(attributes); i++ {
		attr := &attributes[i]
		if attr.Key == attribute {
			attr.Val += " " + value
			return
		}
	}
	newAttribute := html.Attribute{Key: attribute, Val: value}
	(*htmlNode).Raw().Attr = append(attributes, newAttribute)
}

func HandleContainerHTMLChanges(modification Modification, target node.Node) {
	for _, htmlChange := range modification.HTMLChanges {
		tagToChange := QuerySelector(target, htmlChange.Query)
		htmlChangeFn := ReplaceInnerHTMLFromNode
		switch htmlChange.Mode {
		case INNER_HTML:
			{
				htmlChangeFn = ReplaceInnerHTMLFromNode
			}
		case OUTER_HTML:
			{
				htmlChangeFn = ReplaceOuterHTMLFromNode
			}
		case APPEND_HTML:
			{
				htmlChangeFn = PrependHTMLToNode
			}
		case PREPEND_HTML:
			{
				htmlChangeFn = AppendHTMLToNode
			}
		}
		StoreID(target, htmlChange.Query)
		htmlChangeFn(htmlChange.HTML, tagToChange)
	}
}

func StoreID(targetContainer node.Node, firstClass string) {
	id, exists := targetContainer.Attrs().Get("id")
	if exists {
		IDS[firstClass] = id
	}
}

func IsBlank(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

func DeleteNodeFromHTMLTree(deletingNode node.Node) {
	realNode := deletingNode.Raw()
	next := realNode.NextSibling
	prev := realNode.PrevSibling
	realNode.NextSibling.PrevSibling = prev
	realNode.PrevSibling.NextSibling = next
	realNode.NextSibling = nil
	realNode.PrevSibling = nil
	realNode.Parent = nil
}

func NewAttributeChange(query string, mode AttributeMode, attribute string) AttributeChange {
	attrData := strings.SplitN(attribute, "=", 2)

	if len(attrData) < 2 {
		panik(" AttrData is malformed `%s`\nmaybe missing `=`", attribute)
	}

	attributeName := attrData[0]
	attributeValue := strings.ReplaceAll(attrData[1], "\"", "")

	return AttributeChange{
		Query:     query,
		Mode:      mode,
		Attribute: CreateAttribute(attributeName, attributeValue),
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
