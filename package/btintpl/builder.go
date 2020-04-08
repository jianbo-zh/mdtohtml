package btintpl

import (
	"os"
	"text/template"

	"github.com/jianbo-zh/mdtohtml/package/btintpl/htmltpl"
)

// GenArticlePage 生成文章页面
func GenArticlePage(
	dstFilePath string,
	tplFuncMap template.FuncMap,
	title string,
	logoStr string,
	topCates []string,
	curCate string,
	articles []string,
	tags []string,
	content string,
	footer string,
) error {
	file, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	tpl, err := template.New("article").Funcs(tplFuncMap).Parse(htmltpl.ArticleTpl)
	if err != nil {
		return err
	}

	// 生成文章 html
	err = tpl.Execute(file, articleHTMLData{
		Title: title,
		Header: header{
			LogoName: logoStr,
			Cates:    topCates,
		},
		Container: articleContainer{
			Sidebar: sidebar{
				CurCate:  curCate,
				Articles: articles,
				Tags:     tags,
			},
			Content: content,
		},
		Footer: footer,
	})

	return err
}

// GenIndexPage 生成主页
func GenIndexPage(
	dstFilePath string,
	tplFuncMap template.FuncMap,
	title string,
	logoStr string,
	topCates []string,
	allCates []string,
	footer string,
) error {
	file, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	tpl, err := template.New("index").Funcs(tplFuncMap).Parse(htmltpl.IndexTpl)
	if err != nil {
		return err
	}

	// 生成文章 html
	err = tpl.Execute(file, indexHTMLData{
		Title: title,
		Header: header{
			LogoName: logoStr,
			Cates:    topCates,
		},
		Catalogs: allCates,
		Footer:   footer,
	})

	return err
}

// GenTagPage 生成标签页
func GenTagPage(
	dstFilePath string,
	tplFuncMap template.FuncMap,
	title string,
	logoStr string,
	topCates []string,
	curCate string,
	articles []string,
	tags []string,
	tagName string,
	tArticles []string,
	footer string,
) error {
	file, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	tpl, err := template.New("article").Funcs(tplFuncMap).Parse(htmltpl.TagTpl)
	if err != nil {
		return err
	}

	// 生成文章 html
	err = tpl.Execute(file, tagHTMLData{
		Title: title,
		Header: header{
			LogoName: logoStr,
			Cates:    topCates,
		},
		Container: tagContainer{
			Sidebar: sidebar{
				CurCate:  curCate,
				Articles: articles,
				Tags:     tags,
			},
			Content: tagContent{
				CurCate:  curCate,
				TagName:  tagName,
				Articles: tArticles,
			},
		},
		Footer: footer,
	})

	return err
}
