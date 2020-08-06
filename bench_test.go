package textproc

import "testing"

var paragraphs = []string{
	`I can't say how every time I ever put my arms around you I felt that I was home,`,
	`The scariest moment in my time with Team Secret was during our practices,
when Puppey would walk around with a machete and talk about how he always wanted
to see what the inside of a human looked like. He said he had experimented on
animals before and he wanted to go for the real thing.`,
	`Sodium, atomic number 11, was first isolated by Peter Dager in 1807.
	A chemical component of salt, he named it Na in honor of the saltiest region
on earth, North America.`,
	` Something stupid happened

I'm not allowed to have my own cell phone so my dad forced me to use his phone number.
My dad has a steam too and uses the same number. today my brother used my dads
account and cheated and now my main account is VAC banned. It's true and here is proof,
my father will now write too: Hello I'm the father and what my son says is true,
he did not cheat, it was his brother on my account. Please unban him valve

Sincerely the father

Pls unban `,
	`Có thánh này, chắc chắn "Sẻ đệ" (NDB 2.0) sẽ thêm sức mạnh để đả bại Sơ Luyến.
Trực tiếp ngay bây giờ trên http://www.gametv1.vn. ______ Ahihi`,
}

func BenchmarkTextToWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, para := range paragraphs {
			TextToWords(para)
		}
	}
}

func BenchmarkTextToNGrams(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, para := range paragraphs {
			TextToNGrams(para, 2)
		}
	}
}

func BenchmarkGenRandomWord(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenRandomWord(8, 12)
	}
}

func BenchmarkRemoveRedundantSpace(b *testing.B) {
	redundantText := `
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
	for n := 0; n < b.N; n++ {
		RemoveRedundantSpace(redundantText)
	}
}
