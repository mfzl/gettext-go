// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

// SetLocale sets or queries the program's current locale.
func SetLocale(locale string) string {
	return ""
}

// Textdomain sets or retrieves the current message domain.
func Textdomain(domain string) string {
	return ""
}

// BindTextdomain sets directory containing message catalogs.
func BindTextdomain(domain, path string) error {
	return nil
}

// Attempt to translate a text string into the user's native language, by
// looking up the translation in a message catalog.
func Gettext(msgid string) string {
	return PGettext(callerName(2), msgid)
}

// Attempt to translate a text string into the user's native language, by
// looking up the appropriate plural form of the translation in a message
// catalog.
func NGettext(msgid, msgidPlural string, n int) string {
	return PNGettext(callerName(2), msgid, msgidPlural, n)
}

func PGettext(msgctxt, msgid string) string {
	return PNGettext(msgctxt, msgid, "", 0)
}

// Attempt to translate a text string into the user's native language, by
// looking up the appropriate plural form of the translation in a message
// catalog.
func PNGettext(msgctxt, msgid, msgidPlural string, n int) string {
	return msgid
}

// Like Gettext(), but looking up the message in the specified domain.
func DGettext(domain, msgid string) string {
	return DPGettext(domain, callerName(2), msgid)
}

// Like NGettext(), but looking up the message in the specified domain.
func DNGettext(domain, msgid, msgidPlural string, n int) string {
	return DPNGettext(domain, callerName(2), msgid, msgidPlural, n)
}

// Like NGettext(), but looking up the message in the specified domain.
func DPGettext(domain, msgctxt, msgid string) string {
	return DPNGettext(domain, callerName(2), msgid, "", 0)
}

func DPNGettext(domain, msgctxt, msgid, msgidPlural string, n int) string {
	return msgid
}
