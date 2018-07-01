#!/usr/bin/env python
# -*- coding: utf-8 -*-

import webapp2

class MainPage(webapp2.RequestHandler):
    def get(self):
        # input and output
        self.response.headers['Content-Type'] = 'text/html; charset=UTF-8'
        self.response.write(u'こんにちは！')

app = webapp2.WSGIApplication([
    ('/', MainPage),
], debug=True)
