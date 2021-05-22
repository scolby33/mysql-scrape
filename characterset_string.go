// Code generated by "stringer -type=CharacterSet"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[big5-1]
	_ = x[dec8-3]
	_ = x[cp850-4]
	_ = x[hp8-6]
	_ = x[koi8r-7]
	_ = x[latin1-8]
	_ = x[latin2-9]
	_ = x[swe7-10]
	_ = x[ascii-11]
	_ = x[ujis-12]
	_ = x[sjis-13]
	_ = x[hebrew-16]
	_ = x[tis620-18]
	_ = x[euckr-19]
	_ = x[koi8u-22]
	_ = x[gb2312-24]
	_ = x[greek-25]
	_ = x[cp1250-26]
	_ = x[gbk-28]
	_ = x[latin5-30]
	_ = x[armscii8-32]
	_ = x[utf8-33]
	_ = x[ucs2-35]
	_ = x[cp866-36]
	_ = x[keybcs2-37]
	_ = x[macce-38]
	_ = x[macroman-39]
	_ = x[cp852-40]
	_ = x[latin7-41]
	_ = x[cp1251-51]
	_ = x[utf16-54]
	_ = x[utf16le-56]
	_ = x[cp1256-57]
	_ = x[cp1257-59]
	_ = x[utf32-60]
	_ = x[binary-63]
	_ = x[geostd8-92]
	_ = x[cp932-95]
	_ = x[eucjpms-97]
	_ = x[gb18030-248]
	_ = x[utf8mb4-255]
}

const _CharacterSet_name = "big5dec8cp850hp8koi8rlatin1latin2swe7asciiujissjishebrewtis620euckrkoi8ugb2312greekcp1250gbklatin5armscii8utf8ucs2cp866keybcs2maccemacromancp852latin7cp1251utf16utf16lecp1256cp1257utf32binarygeostd8cp932eucjpmsgb18030utf8mb4"

var _CharacterSet_map = map[CharacterSet]string{
	1:   _CharacterSet_name[0:4],
	3:   _CharacterSet_name[4:8],
	4:   _CharacterSet_name[8:13],
	6:   _CharacterSet_name[13:16],
	7:   _CharacterSet_name[16:21],
	8:   _CharacterSet_name[21:27],
	9:   _CharacterSet_name[27:33],
	10:  _CharacterSet_name[33:37],
	11:  _CharacterSet_name[37:42],
	12:  _CharacterSet_name[42:46],
	13:  _CharacterSet_name[46:50],
	16:  _CharacterSet_name[50:56],
	18:  _CharacterSet_name[56:62],
	19:  _CharacterSet_name[62:67],
	22:  _CharacterSet_name[67:72],
	24:  _CharacterSet_name[72:78],
	25:  _CharacterSet_name[78:83],
	26:  _CharacterSet_name[83:89],
	28:  _CharacterSet_name[89:92],
	30:  _CharacterSet_name[92:98],
	32:  _CharacterSet_name[98:106],
	33:  _CharacterSet_name[106:110],
	35:  _CharacterSet_name[110:114],
	36:  _CharacterSet_name[114:119],
	37:  _CharacterSet_name[119:126],
	38:  _CharacterSet_name[126:131],
	39:  _CharacterSet_name[131:139],
	40:  _CharacterSet_name[139:144],
	41:  _CharacterSet_name[144:150],
	51:  _CharacterSet_name[150:156],
	54:  _CharacterSet_name[156:161],
	56:  _CharacterSet_name[161:168],
	57:  _CharacterSet_name[168:174],
	59:  _CharacterSet_name[174:180],
	60:  _CharacterSet_name[180:185],
	63:  _CharacterSet_name[185:191],
	92:  _CharacterSet_name[191:198],
	95:  _CharacterSet_name[198:203],
	97:  _CharacterSet_name[203:210],
	248: _CharacterSet_name[210:217],
	255: _CharacterSet_name[217:224],
}

func (i CharacterSet) String() string {
	if str, ok := _CharacterSet_map[i]; ok {
		return str
	}
	return "CharacterSet(" + strconv.FormatInt(int64(i), 10) + ")"
}