package gurl

/*
#include <stdio.h>
#include <unistd.h>
#include <math.h>
#include <curl/curl.h>

int progress_callback_old( char *Bar,double t,double d,double ultotal,double ulnow)
{
    if ( d != 0 ) {
    fprintf(stderr, "\r%i",(int)(d/t*100));
    //usleep(59000);
    fflush(stderr);
    }
    return 0;
}

int progress_callback(void* ptr, double dtotal, double dnow, 
                    double ultotal, double ulnow)
{
    // how wide you want the progress meter to be
    int width =79;
    double frac = dtotal / dtotal;
    // part of the progressmeter that's already "full"
    int dotz = frac * width;
    // create the "meter"
    int ii=0;
    // part  that's full already
    for ( ; ii < dotz;ii++) {
        printf("#");
    }
    // remaining part (spaces)
    for ( ; ii < width;ii++) {
        printf(" ");
    }
    // and back to line begin - do not forget the fflush to avoid output buffering problems!
    printf(" %2.0f%%\r",frac*100);
    //printf("\r");
    fflush(stdout);
    return 0;
}

void download(char* url, char* dest) {
    fprintf(stderr, "%s", url);
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
        curl_easy_setopt(curl, CURLOPT_PROGRESSFUNCTION, progress_callback);
        //curl_easy_setopt(curl, CURLOPT_PROGRESSDATA, &progressbar);
        res = curl_easy_perform(curl);
        curl_easy_cleanup(curl);
        fclose(fp);
    }
    printf("\n");
    fflush(stdout);
}
*/
import "C"
import "path"

func Download(url string, dest string) {
	_, file := path.Split(url)
	file = dest + file
	C.download(C.CString(url), C.CString(file))
}

func Version() string {
	version := C.GoString(C.curl_version())
	return version
}
