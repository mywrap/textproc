# Text processing

Extract information from text and HTML.  
Created because of [daominah/scraper](
https://github.com/daominah/scraper), a scraper that can filter nearly
duplicate news and extract keywords.

## Functions

* **NormalizeText** normalize different representations of a character.
* **TextToNGrams** creates a set of n-gram (lowercase) from input text.

* **HTMLXPath** finds all html nodes match the xpath query.
* **HTMLGetHREFs** returns all URLs (absolute form) in a HTML.
* **HTMLGetText** get content from a HTML (javascript, spaces removed)

## Example
Detail in [text_test.go](./text_test.go) and 
[html_test.go](./html_test.go).
