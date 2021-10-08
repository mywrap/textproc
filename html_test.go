package textproc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
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

func TestHTMLGetTextF1(t *testing.T) {
	file, err := ioutil.ReadFile("html_test_file1.html")
	if err != nil {
		t.Fatal(err)
	}
	htmlTree, err := html.Parse(bytes.NewReader(file))
	if err != nil {
		t.Fatalf("error html Parse: %v", err)
	}
	text := HTMLGetText(htmlTree)
	if expected := `Google
Gmail
Hình ảnh
Đăng nhập
Xóa
Báo cáo các gợi ý không phù hợp
Google có các thứ tiếng:
English
Français
中文(繁體)
Việt Nam
Bảo mật
Điều khoản
Cài đặt
Cài đặt tìm kiếm
Tìm kiếm nâng cao
Dữ liệu của bạn trong Tìm kiếm
Hoạt động tìm kiếm
Trợ giúp tìm kiếm
Gửi phản hồi
Quảng cáo
Doanh nghiệp
Giới thiệu
Cách hoạt động của Tìm kiếm`; text != expected {
		t.Error("unexpected HTMLGetText file1:")
		fmt.Println(text)
		//aLines := strings.Split(text, "\n")
		//eLines := strings.Split(text, "\n")
		//for i, _ := range aLines {
		//	t.Log(aLines[i] == eLines[i])
		//}
	}
}

func TestHTMLGetTextF2(t *testing.T) {
	file, err := ioutil.ReadFile("html_test_file2.html")
	if err != nil {
		t.Fatal(err)
	}
	htmlTree, err := html.Parse(bytes.NewReader(file))
	if err != nil {
		t.Fatalf("error html Parse: %v", err)
	}
	text := HTMLGetText(htmlTree)
	if text != `Hàng hóa
Năng lượng
Thứ năm, 6/8/2020, 07:10 (GMT+7)
Tồn kho tại Mỹ giảm, giá dầu lên đỉnh 5 tháng, vàng tiếp tục lập đỉnh lịch sử
Chốt phiên 5/8, giá dầu Brent, WTI đều tăng.
Giá vàng tiếp tục tăng, lập đỉnh lịch sử mới nhờ USD suy yếu và lợi nhuận từ trái phiếu chính phủ Mỹ giảm.
Như Tâm (Theo Reuters)
Thứ năm, 6/8/2020, 07:10 (GMT+7)
Giá dầu Brent tương lai tăng 74 cent, tương đương 1,7%, lên 45,17 USD/thùng.
Giá dầu WTI tương lai tăng 49 cent, tương đương 1,2%, lên 42,19 USD/thùng.
Đầu phiên 5/8, giá của hai loại dầu đều có lúc tăng hơn 4%.
Tồn kho dầu thô tại Mỹ trong tuần kết thúc ngày 31/7 giảm 7,4 triệu thùng, theo Cơ quan thông tin năng lượng Mỹ (EIA), vượt dự báo giảm 3 triệu thùng từ giới phân tích.
USD suy yếu, khiến dầu rẻ hơn với những người nắm giữ đồng tiền khác, cũng hỗ trợ thị trường.
“Thị trường hàng hóa hưởng lợi từ USD suy yếu và dầu chắc chắn cũng có phần”, theo Craig Erlam, nhà phân tích tại OANDA, New York.
Thị trường năng lượng còn được thúc đẩy bởi các dấu hiệu cho thấy quá trình thương lượng giữa Nhà Trắng và phe Dân chủ tại quốc hội về gói hỗ trợ kinh tế tiếp theo đang có tiến triển, dù hai bên chưa có sự đồng thuận.
Số liệu cho thấy hoạt động tại các nhà máy ở Mỹ trong tuần có cải thiện, được giới phân tích coi là dấu hiệu phục hồi kinh tế. Hoạt động kinh doanh tại eurozone cũng dần đi lên khi một số quy định hạn chế để ngăn Covid-19 lây lan được nới lỏng.
Số ca nhiễm Covid-19 trên thế giới vẫn tăng, đe dọa đà phục hồi trong lực cầu năng lượng. Thế giới ghi nhận hơn 700.000 trường hợp tử vong vì đại dịch tính đến ngày 5/8 với Mỹ, Brazil, Ấn Độ và Mexico có số người chết cao nhất.
“Nhu cầu xăng có thể giảm 7% trong quý III so với cùng kỳ năm trước, cho thấy đà phục hồi tiếp tục chững lại và khả năng quay trở về mức như năm 2019 ngay trong năm nay khó xảy ra”, JBC Energy nhận định. Nhu cầu nhiên liệu máy bay dự báo giảm 50% trong quý III so với cùng kỳ năm trước.
Kim loại quý
Giá vàng ngày 5/8 tiếp tục tăng, lập đỉnh lịch sử mới nhờ USD suy yếu và lợi nhuận từ trái phiếu chính phủ Mỹ giảm khiến nhà đầu tư chuyển hướng sang kim loại quý này. Giá vàng đã tăng 34% kể từ đầu năm nay.
Giá vàng giao ngay tại sàn New York tăng 20,1 USD lên 2.039,5 USD/ounce, trong phiên có lúc chạm 2.055,1 USD/ounce.
Giá vàng tương lai tăng 1,4% lên 2.049,3 USD/ounce, trong phiên có lúc chạm 2.070,3 USD/ounce.
Giá vàng giao ngay tại sàn New York ngày 5/8.
Giá bạc tăng 4,3% lên 27,13 USD/ounce, cao nhất kể từ tháng 4/2013, và đã tăng 50% kể từ đầu năm.
Giá platinum tăng 2,7% lên 962,63 USD/ounce.
giá dầu thế giới
giá vàng thế giới
Brent
WTI` {
		t.Errorf("unexpected HTMLGetText:")
		fmt.Println(text)
	}
}

func TestHTMLParseToNode(t *testing.T) {
	resp, err := http.Get("https://google.com/should-be-not-found")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	node := HTMLParseToNode(resp.Body) // ignore readAll body error
	if node == nil {
		t.Fatal("unexpected nil return")
	}
}

func TestHTMLParseToNode2(t *testing.T) {
	resp, err := http.Get("https://github.com/should-be-not-found")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	textInBody := HTMLGetText(HTMLParseToNode(body))
	t.Logf("len respBody: %v, len textInBody: %v", len(body), len(textInBody))
	t.Logf("textInBody: %v", textInBody)
	if len(textInBody) < 1 || len(textInBody) > 4095 {
		t.Errorf("too long textInBody: %v", len(textInBody))
	}
}
