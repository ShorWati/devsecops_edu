from http.server import SimpleHTTPRequestHandler, HTTPServer

class MyHandler(SimpleHTTPRequestHandler):
	def do_GET(self):
		self.send_response(200)
		self.send_header('Content-type', 'text/html; charset=utf-8')
		self.end_headers()
		self.wfile.write(b"<h1>test!</h1>")

print("Server started, ports: 8000...")
HTTPServer(('0.0.0.0', 8000),MyHandler).serve_forever()
