package lb

import (
	api "github.com/mats9693/listen_bilibili/api/go"
	"net/http"
	"os"
)

type handlerFunc = func(r *http.Request) ([]byte, error)

type Handler struct {
	handlerFuncs map[string]handlerFunc // request uri - func(*http.request) ([]byte, error)
}

var handlerIns = &Handler{
	handlerFuncs: make(map[string]handlerFunc, 3),
}

func GetHandler() *Handler {
	if len(handlerIns.handlerFuncs) < 1 {
		initHandler()
	}

	return handlerIns
}

func initHandler() {
	handlerIns.HandleFunc(api.URI_GetList, onGetList)
	handlerIns.HandleFunc(api.URI_GetOriginURL, onGetOriginURL)
	handlerIns.HandleFunc(api.URI_ModifyList, onModifyList)

	Println("> Init HTTP Handler Finished.")
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// allow cross-origin
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	// skip 'http options' request
	if request.Method == http.MethodOptions {
		response(writer, []byte(""))
		return
	}

	_ = request.ParseMultipartForm(32 << 20) // 32 MB
	params := request.PostForm.Encode()
	if len(params) > 1024 {
		params = "too lang"
	}

	// log req
	Printf("> Receive new request. uri: %s, params: %s\n", request.RequestURI, params)

	// invoke handleFunc func
	var res []byte
	handlerFuncIns, ok := h.handlerFuncs[request.RequestURI]
	if !ok { // unknown uri, regard as web source, like 'xxx.js'
		dir, _ := os.Getwd()
		path := dir + "/ui/dist" + request.RequestURI
		if request.RequestURI == "/" { // bind html file
			path += "index.html"
		}
		http.ServeFile(writer, request, path)
		return
	}

	res, err := handlerFuncIns(request)
	if err != nil {
		res = []byte(err.Error())
	}

	// log res
	Printf("> Handle request %s success: %t\n", request.RequestURI, err == nil)

	// response
	response(writer, res)
}

func (h *Handler) HandleFunc(pattern string, hf handlerFunc) {
	Println("> register http handler on uri: ", pattern)

	h.handlerFuncs[pattern] = hf
}

func response(writer http.ResponseWriter, data []byte) {
	_, err := writer.Write(data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
