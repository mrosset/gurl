package gurl

/*
#include <curl/curl.h>
#include <curl/types.h>
#include <curl/easy.h>

size_t write_data(void *ptr, size_t size, size_t nmemb, FILE *stream) {
    size_t written = fwrite(ptr, size, nmemb, stream);
    return written;
}

void download(char* url, char* dest) {
    CURL *curl;
    FILE *fp;
    CURLcode res;
    curl = curl_easy_init();
    if(curl) {
        fp = fopen(dest,"wb");
        curl_easy_setopt(curl, CURLOPT_URL, url);
        //curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_data);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, fp);
        curl_easy_setopt(curl, CURLOPT_VERBOSE, 1);
        curl_easy_setopt(curl, CURLOPT_NOPROGRESS, 0);
        res = curl_easy_perform(curl);
        curl_easy_cleanup(curl);
        fclose(fp);
    }
}
*/
import "C"
import "path"

func Download(url string) {
	_, file := path.Split(url)
	C.download(C.CString(url), C.CString("./"+file))
	println(url)
}

func Version() string {
	version := C.GoString(C.curl_version())
	return version
}
