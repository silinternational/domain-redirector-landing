# Simple webserver / redirector for use when a domain changes

Sometimes hostnames need to change. In the past we've run automatic redirector services
that redirect from an old domain to a new one, carrying along the path so it is transparent
to the end user. We've learned though that by doing this end users never realize something 
changed and then continue to depend on the old domain and the redirector to access stuff.
Eventually we want to shut down these redirectors or let go of old domains.

This new redirector service presents a landing page instead of just redirecting. This 
way the end user is forced to acknowledge the url has changed and hopefully they'll 
update their bookmarks or links accordingly. 

## Configuration
This service requires one environment variable, but supports a few more for additional
content.

__REQUIRED:__

 - `NEW_HOST` - The new hostname. The URL generated for the user will be formatted as 
   `https://{NEW_HOST}{REQUEST_URI}`. The _https_ is hard coded to prevent redirecting
   to an http url, so only provide the hostname. 

_OPTIONAL:_

 - `ADDITIONAL_MESSAGE` - By default if this is not provided, the landing page will have a
   message that says `<em>{{.OldHost}}</em> has changed to <em>{{.NewHost}}</em>`. Use
   `ADDITIONAL_MESSAGE` to override this message. `ADDITIONAL_MESSAGE` is considered HTML 
   safe, so you can include HTML in this message and it will not be escaped. 
   
 - `REDIRECT_END_DATE` - A human readable date that you intend to stop this redirector service.
   This is intended to instill a sense of inevitability so that end users do not perpetually 
   depend on this redirector service. 
   
 - `MORE_INFO_URL` - If this URL is provided, a message is added to the page that says 
   `<p>Learn more about this change at: <a target="_blank" href="{{.MoreInfoURL}}">{{.MoreInfoURL}}</a></p>`
   
 - `SKIP_LANDING_PAGE` - If set to `true` users will be automatically redirected without seeing the landing page
   
## Running Locally
Just run `docker-compose up -d`. Feel free to edit environment variables in the `docker-compose.yml` file
first if you like. 


## License
MIT License

Copyright (c) 2020 SIL International

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

