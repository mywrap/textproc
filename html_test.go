package textproc

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestHtmlUtils(t *testing.T) {
	htmlStr := `
<html>

<head><title>A Simple HTML Document</title></head>

<body>
    <div id="main">
        The main div
        <p>This is a very simple HTML document</p>
        <p>It only has two paragraphs</p>
        <a href="https://www.w3schools.com">This is a link</a>
        <a href="https://www.google.com.vn#frag2" id="xxx2">Another link</a>
        <a href="/imghp" class="c1">Link3</a>
        <a href="https://www.google.com.vn#frag3" class="c2">Link4</a>
        <img src="/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png" height=50px>
        
        <div id="dev_page_content_wrap" class=" ">
            <h3>This is dev_page_content_wrap_h3</h3>
            <a href="https://www.google.com.vn" class="c1">Link5</a>
            <a href="https://www.google.com.vn" class="c2">Link6</a>
            <img class="rg_i Q4LuWd" src="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEASABIAAD/2wBDABALDA4MChAODQ4SERATGCgaGBYWGDEjJR0oOjM9PDkzODdASFxOQERXRTc4UG1RV19iZ2hnPk1xeXBkeFxlZ2P/2wBDARESEhgVGC8aGi9jQjhCY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2P/wgARCAAyADgDAREAAhEBAxEB/8QAGgAAAgMBAQAAAAAAAAAAAAAAAwUAAgQBBv/EABcBAQEBAQAAAAAAAAAAAAAAAAEAAgP/2gAMAwEAAhADEAAAAbVWmQ0GOQ1lpFW6n+d5DWgB6zl0KYZy4zoQlQMYWUOGcus74XGlRvNa5lFvjfa5Q2EO0jCtYaSiDWmUVkBAYhEWxf/EAB8QAAIBAwUBAAAAAAAAAAAAAAECAAMREgQQEyEiMf/aAAgBAQABBQIV9OYaq5YJGFMHjS701EZV2pegzY0yLmoMlY3Rj58TRBS9yJxiKOqvx0bHCaPp9iZe8zA2p3yVyIWM7mTEm8XT01jL091nMZyO0CGFGEvCY2wgjfP/xAAbEQACAgMBAAAAAAAAAAAAAAABIAAQERIwIf/aAAgBAwEBPwGsTXibDhwoniYocP/EABoRAAIDAQEAAAAAAAAAAAAAAAARARAgMAL/2gAIAQIBAT8BpjGPUwecPrOHT0+H/8QAJBAAAgEEAQIHAAAAAAAAAAAAAAERAhAxMpEhMBIiM0FCYXH/2gAIAQEABj8CwuD06Y/DWng1p4I8NJquDCskYdp9yljPkYId4S7Xlm+CE0jqzW3Q1ZiLfXY//8QAIRABAAICAgMAAwEAAAAAAAAAAQARITFRYRBBcYGRobH/2gAIAQEAAT8hcqv7Iyd2MtP6kCGzpLKLfhE6ioP+Xi25bqCmVGtblO3puYRp/Yqm4n4JXGENLHMyEuKWmJSLhoAgk1uW5hHF2RLgz0yn1PZfFop9BBsl/GK9JC0LlL5zKhyTTFe8wMTB1NLInJB0ION1O4i3LE5fB3NY3mJvmf/aAAwDAQACAAMAAAAQWT/cSL40N/jN645Z50AFTAd83//EABoRAAMBAAMAAAAAAAAAAAAAAAABERAgITH/2gAIAQMBAT8QxM0JoRjXBdPE4L3wRMo2ekhcLIeDVxHSZc9EQTGiCEE7ELl//8QAGxEAAwEBAQEBAAAAAAAAAAAAAAEREDEhMEH/2gAIAQIBAT8QqKOMwKh6/SwoqNUPGNkoQSOFpBoeU6J4ySLvg6qNn+4mUbKcfD//xAAgEAEAAwADAQADAQEAAAAAAAABABEhMUFRkXHB0RCh/9oACAEBAAE/ECQY+B+oONVJN16QIhDtn8oP2HVpU/I0wfFNGD2hDOHQ8H/E6taXuLSbuvA0+y36nI2IA52J0jdXsX2XG9tKDk+RwvmxfcAIBzfJyjO0cQOMDqYknhLsiQqdO4L1+TfzgewuLcoIAUdvYPBqmk0USw5QlvUXqED2ADP3/SGWXONJqVP5QI0eJoYr5vizuPCc72/7LUzwFENw62kcRieGREK1Vb/JuEnvuV/JNq8hG390o/ujd1gHQTGM/EscvsZu49n/2Q==" data-deferred="1" jsname="Q4LuWd" alt="DD 둔둔이 | 사랑스러운 새끼 고양이, 새끼 고양이, 고양이">
        </div>
    </div>
    
    <p id="demo"></p>
    <script>
        document.getElementById("demo").innerHTML = "Hello JavaScript!";
    </script>
</body>

</html>`
	baseURL := "https://www.google.com.vn"

	root, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		t.Fatalf("error html Parse: %v", err)
	}

	elems, err := HTMLXPath(root, `//a[@class='c1']`)
	if err != nil || len(elems) != 2 {
		t.Fatal(err)
	}

	if HTMLGetText(elems[0]) != "Link3" ||
		HTMLGetText(elems[1]) != "Link5" {
		t.Error()
	}

	urls := HTMLGetHREFs(baseURL, root)
	if len(urls) != 3 ||
		urls[0] != "https://www.google.com.vn" ||
		urls[1] != "https://www.google.com.vn/imghp" ||
		urls[2] != "https://www.w3schools.com" {
		t.Errorf("error HtmlGetHrefs: real: %v", urls)
	}

	text := HTMLGetText(root)
	if strings.Contains(text, "getElementById") {
		t.Errorf("error javascript in text: %v", text)
	}

	imgNodes, _ := HTMLXPath(root, `//img`)
	if len(imgNodes) != 2 {
		t.Fatalf("error nImgs: real: %v, expected: 2", len(imgNodes))
	}
	src0 := HTMLGetImgSrc(baseURL, imgNodes[0])
	e := "https://www.google.com.vn" +
		"/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	if src0 != e {
		t.Errorf("error HtmlGetImgSrc: real: %v, expected: %v", src0, e)
	}
	src1 := HTMLGetImgSrc(baseURL, imgNodes[1])
	if src1 != "" {
		t.Errorf("real: %v, expected empty string", src1)
	}

	_, err = HTMLXPath(root, `//a[@class='c1'invalidExample`)
	if err == nil {
		t.Error("expect error invalid xpath")
	}
}
