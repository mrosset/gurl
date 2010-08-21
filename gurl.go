package gurl

/*
#include <curl/curl.h>

void download(char* url, char* dest) {
    CURL *curl;
    FILE *fp;
    CURLcode res;
    curl = curl_easy_init();
    if(curl) {
        fp = fopen(dest,"wb");
        curl_easy_setopt(curl, CURLOPT_URL, url);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, fp);
        curl_easy_setopt(curl, CURLOPT_VERBOSE, 0);
        curl_easy_setopt(curl, CURLOPT_NOPROGRESS, 0);
        res = curl_easy_perform(curl);
        curl_easy_cleanup(curl);
        fclose(fp);
    }
}
*/
import "C"
import "path"

func Download(url string, dest string) {
	_, file := path.Split(url)
	file = dest + file
	C.download(C.CString(url), C.CString(file))
	println(url)
}

func Version() string {
	version := C.GoString(C.curl_version())
	return version
}
