package utilities

import (
	"os"
	"testing"
)

func TestGetFileNameFromURL(t *testing.T) {
	url := "https://s3.us-west-2.amazonaws.com/secure.notion-static.com/eab97ee2-36b0-4c75-a3b6-547fd6271b0d/Screen_Shot_2022-11-15_at_2.14.09_PM.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20221115%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20221115T222215Z&X-Amz-Expires=3600&X-Amz-Signature=2d33ebb17d788eb2561b87ba952ed370acdbe63de13f4df1e22ef1c937d5977c&X-Amz-SignedHeaders=host&x-id=GetObject"
	expected := "Screen_Shot_2022-11-15_at_2.14.09_PM.png"
	actual, _ := GetFileNameFromURL(url)

	if expected != actual {
		t.Fatalf("TestGetFileNameFromURL failed \nexpected: %s \ngot: %s ", expected, actual)
	} else {
		t.Logf("TestGetFileNameFromURL success \nexpected: %s \ngot: %s ", expected, actual)
	}

}

func TestGetAssetPath(t *testing.T) {
	os.Setenv("ASSET_PATH", "assets/img/posts/$PAGE_URI$/$FILE_NAME$")
	file := "thebunny.mp4"
	pageUri := "a-talk-about-the-bunny"

	expected := "assets/img/posts/a-talk-about-the-bunny/thebunny.mp4"
	actual := GetAssetPath(file, pageUri)

	if expected != actual {
		t.Fatalf("TestGetAssetPath failed \nexpected: %s \ngot: %s", expected, actual)
	} else {
		t.Logf("TestGetAssetPath success \nexpected: %s \ngot: %s", expected, actual)
	}
}
