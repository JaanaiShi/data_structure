package main

import (
	"fmt"
	"log"
	"net/url"

	git "github.com/libgit2/git2go/v33"
)

func Clone(addr, path string) {
	r, err := git.Clone(addr, path, &git.CloneOptions{})
	if err != nil {
		log.Fatal(err)
	}
	BranchIterator, err := r.NewBranchIterator(git.BranchLocal)
	if err != nil {
		log.Println(err)
	}
	log.Println("BranchIterator is ", BranchIterator)

	var branchs []map[string]interface{}
	//函数原型   type BranchIteratorFunc func(*Branch, BranchType) error
	f := func(b *git.Branch, t git.BranchType) error {
		branch := make(map[string]interface{})
		name, err := b.Name()
		if err != nil {
			return err
		}
		branch["name"] = name
		branch["head"] = false
		head, err := b.IsHead()
		if err != nil {
			return err
		}
		if head == true {
			branch["head"] = true
		}
		branchs = append(branchs, branch)
		return nil
	}

	err = BranchIterator.ForEach(f)

	if err != nil {
		log.Println("err is ", err)
	}
	log.Println("Name is", branchs)
}

func main() {
	addr := "https://shi18945010548@dev.azure.com/shi18945010548/note/_git/note/git-upload-pack"
	username := "shi18945010548"
	password := "etqwrq3ubxz4gdbyy7zig42hardkantkhumyurwipx6dsscpklqa"
	urlObj, err := url.Parse(addr)
	if err != nil {
		log.Fatal("解析url失败", err)
	}

	fmt.Println(urlObj.User)
	urlObj.User = url.UserPassword(username, password)
	addr = urlObj.String()
	fmt.Println("new addr", addr)

	Clone(addr, "./")
}
