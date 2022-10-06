# noteserver-go
Show my markdown html notes with go web server  
I will exercise my go web program ability with this project

## 2022/10/6
1. Find `http.FileServer` can show the simple file directories and show the contents of the files without parsing the html files  
So I've commented out the parsing code for now, leaving it for later optimisation  
2. Visiting the page will generate multiple `GET /favicon.ico HTTP/1.1`  
   Temporarily use `return` to process the request

