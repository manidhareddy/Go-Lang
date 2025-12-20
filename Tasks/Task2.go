package main
//https://pkg.go.dev/golang.org/x/net@v0.48.0/html
import (
	"fmt"
	"strings"
	"bytes"
	"golang.org/x/net/html"
	"os"
)

func count (doc *html.Node, words *int,pics *int){
	if doc.Type == html.TextNode  {
		*words += len(strings.Fields(doc.Data))
} else  {
	if doc.Type == html.ElementNode && doc.Data == "img"{
		*pics += 1
	}
}
	for node := doc.FirstChild;node!=nil;node = node.NextSibling {
		count(node,words,pics)
	} 
}

func countWordsAndPics(doc *html.Node) (int,int){
	var words, pics int
	count(doc,&words,&pics)

	return words, pics

}
var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
  </body>
</html>`

func main(){
	doc,err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil{
		fmt.Fprintf(os.Stderr,"unable to parse html %s\n",err)
		os.Exit(-1)
	}

	word, pics := countWordsAndPics(doc)

	fmt.Printf("%d words and %d images are present \n",word,pics)
}