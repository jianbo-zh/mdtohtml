package btintpl

type header struct {
	LogoName string
	Cates    []string
}

type sidebar struct {
	CurCate  string
	Articles []string
	Tags     []string
}

type articleContainer struct {
	Sidebar sidebar
	Content string
}

type tagContent struct {
	CurCate  string
	TagName  string
	Articles []string
}

type tagContainer struct {
	Sidebar sidebar
	Content tagContent
}

type articleHTMLData struct {
	Title     string
	Header    header
	Container articleContainer
	Footer    string
}

type tagHTMLData struct {
	Title     string
	Header    header
	Container tagContainer
	Footer    string
}

type indexHTMLData struct {
	Title    string
	Header   header
	Catalogs []string
	Footer   string
}
