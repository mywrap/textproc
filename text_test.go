package textproc

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

func TestTextInit(t *testing.T) {
	// t.Logf("Numeric: %#v", Numeric)
	// t.Logf("LowerAlpha: %v", LowerAlpha)
	// t.Logf("UpperAlpha: %v", UpperAlpha)
	// t.Logf("AlphaNumeric: %v", AlphaNumeric)
	// t.Logf("AlphaNumericList: %v", AlphaNumericList)
}

func TestTextToNGrams(t *testing.T) {
	text := `Có thánh này, chắc chắn "Sẻ đệ" (NDB 2.0) sẽ thêm sức mạnh để đả bại Sơ Luyến.
Trực tiếp ngay bây giờ trên http://www.gametv1.vn. ______ Ahihi`
	words := TextToWords(text)
	jbs, err := json.Marshal(words)
	if err != nil {
		t.Error(err)
	}
	if string(jbs) != `["Có","thánh","này","chắc","chắn","Sẻ","đệ","NDB","2.0","sẽ","thêm","sức","mạnh","để","đả","bại","Sơ","Luyến","Trực","tiếp","ngay","bây","giờ","trên","http://www.gametv1.vn","Ahihi"]` {
		t.Error(string(jbs))
	}

	nGrams := TextToNGrams(text, 3)
	jbs, err = json.Marshal(nGrams)
	if err != nil {
		t.Error(err)
	}
	if string(jbs) != `{"2.0 sẽ thêm":1,"bây giờ trên":1,"bại sơ luyến":1,"chắc chắn sẻ":1,"chắn sẻ đệ":1,"có thánh này":1,"giờ trên http://www.gametv1.vn":1,"luyến trực tiếp":1,"mạnh để đả":1,"ndb 2.0 sẽ":1,"ngay bây giờ":1,"này chắc chắn":1,"sơ luyến trực":1,"sẻ đệ ndb":1,"sẽ thêm sức":1,"sức mạnh để":1,"thánh này chắc":1,"thêm sức mạnh":1,"tiếp ngay bây":1,"trên http://www.gametv1.vn ahihi":1,"trực tiếp ngay":1,"đả bại sơ":1,"để đả bại":1,"đệ ndb 2.0":1}` {
		t.Error(string(jbs))
	}
}

func TestHashTextToInt64(t *testing.T) {
	nWords := 1000000 // fast
	words := make(map[string]bool)
	hashes := make(map[int64]bool)
	for i := 0; i < nWords; i++ {
		word := GenRandomWord(1, 4, AlphaNumericList)
		words[word] = true
		hashes[HashTextToInt(word)] = true
		hashes[HashTextToInt(word)] = true
	}
	if math.Abs(float64(len(words)-len(hashes))) > 1 { // unique
		t.Error()
	}
}

func TestTextNormalize(t *testing.T) {
	in := `VE bị phạt 100 triệu đồng do không công bố thông tin đúng quy định
Ngày 30/09, Thanh tra Ủy ban Chứng khoán Nhà nước (UBCKNN) đã quyết định xử phạt
vi phạm hành chính trong lĩnh vực chứng khoán và thị trường chứng khoán đối với 
Tổng Công ty Tư vấn thiết kế dầu khí - CTCP (HNX:PVE). 
Cụ thể, Công ty này đã không công bố thông tin tài liệu.`
	out := NormalizeText(in)
	if out != `VE bị phạt 100 triệu đồng do không công bố thông tin đúng quy định
Ngày 30/09, Thanh tra Ủy ban Chứng khoán Nhà nước (UBCKNN) đã quyết định xử phạt
vi phạm hành chính trong lĩnh vực chứng khoán và thị trường chứng khoán đối với 
Tổng Công ty Tư vấn thiết kế dầu khí - CTCP (HNX:PVE). 
Cụ thể, Công ty này đã không công bố thông tin tài liệu.` {
		t.Error(out)
	}
}

func TestRemoveRedundantSpace(t *testing.T) {
	input := `
Google
Gmail
Hình ảnh
Đăng nhập		
	
Xóa


Báo cáo các gợi ý không phù hợp
Google có các thứ tiếng:  
English
    
Français
    
中文（繁體）
  
Việt Nam
Giới thiệu
  Cách hoạt động của Tìm kiếm  `
	e := `Google
Gmail
Hình ảnh
Đăng nhập
Xóa
Báo cáo các gợi ý không phù hợp
Google có các thứ tiếng:
English
Français
中文（繁體）
Việt Nam
Giới thiệu
Cách hoạt động của Tìm kiếm`
	if r := RemoveRedundantSpace(input); r != e {
		t.Errorf("unexpected HTMLGetText:")
		fmt.Println(r)
	}
}

func TestGenRandomWord(t *testing.T) {
	for i := 0; i < 5; i++ {
		word := GenRandomWord(8, 12, AlphaNumericEnList)
		// t.Logf("len: %2d, lenBytes: %2d, word: %v", len([]rune(word)), len([]byte(word)), word)
		if !(8 <= len([]rune(word)) && len([]rune(word)) <= 12) {
			t.Error("wrong GenRandomWord len")
		}
	}
}

func TestGenRandomVarName(t *testing.T) {
	for i := 0; i < 5; i++ {
		word := GenRandomVarName(8)
		t.Logf("len: %2d, lenBytes: %2d, word: %v", len([]rune(word)), len([]byte(word)), word)
	}
}

func TestRemoveVietnamDiacritic(t *testing.T) {
	// d1 := "Đ"
	// d2 := "Ð"
	// t.Logf("%v %v %v %v %v",
	//	len([]rune(d1)), d1, len([]rune(d2)), d2, d1 == d2)

	// t.Logf("removeVietnamDiacritic %c", removeVietnamDiacritic('đ'))

	for _, test := range []struct {
		in  string
		out string
	}{
		{in: "Đào", out: "Dao"},
		{in: "NGUYỄN NGỌC THUẬN", out: "NGUYEN NGOC THUAN"},
		{in: "Hải Ðường", out: "Hai Duong"},
		{in: "office", out: "office"},
		{in: "đ", out: "d"},
	} {
		r, e := RemoveVietnamDiacritic(test.in), test.out
		if r != e {
			t.Errorf("error RemoveVietnamDiacritic: real: %v, expected: %v", r, e)
		}
	}
}
