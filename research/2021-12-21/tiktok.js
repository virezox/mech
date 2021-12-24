const request = require("request")
const url = require("url")
const fs = require("fs")
const randomMac = require("random-mac")
const uuid = require("uuid-random")
const xorCrypt = require("xor-crypt")
const formurlencoded = require("form-urlencoded").default

const TikTok = function() {
	this.session = {}
	this.device = {}
}

TikTok.prototype.generateDevice = async function() {
	const randDevice = getRandFromFile("devices.json")
	const randCarrier = getRandFromFile("carriers.json")
	this.device = {
		default: {
			ac: "wifi",
			channel: "googleplay",
			aid: "1233",
			app_name: "musical_ly",
			version_code: "130211",
			version_name: "13.2.11",
			device_platform: "android",
			ab_version: "13.2.11",
			ssmix: "a",
			device_type: randDevice[1],
			device_brand: randDevice[0],
			language: "en",
			os_api: "25",
			os_version: "7.1.2",
			uuid: randomString(15, "#"),
			openudid: randomString(15, "a#"),
			manifest_version_code: "2019092901",
			resolution: randDevice[2],
			dpi: randDevice[3],
			update_version_code: "2019092901",
			app_type: "normal",
			sys_region: "US",
			is_my_cn: "0",
			"pass-route": "1",
			mcc_mnc: randCarrier[0] + randCarrier[1],
			"pass-region": "1",
			timezone_name: "America/New_York",
			carrier_region_v2: randCarrier[0],
			timezone_offset: "0",
			build_number: "13.2.11",
			region: "US",
			uoo: "0",
			app_language: "en",
			carrier_region: "US",
			locale: "en",
			ac2: "wifi5g"
		},
		"user-agent":
			"com.zhiliaoapp.musically/2019092901 (Linux; U; Android 7.1.2 en; " + randDevice[1] + "; Build/" + randDevice[1] + "; Cronet/58.0.2991.0)",
		mac: randomMac(),
		google_aid: uuid(),
		clientudid: uuid(),
		carrier: randCarrier
	}
}

TikTok.prototype.registerDevice = async function() {
	const protocol = "https"
	const hostname = "log2.musical.ly"
	const path = "service/2/device_register/"
	const requestUrl = buildUrl(protocol, hostname, path, this.device.default)
	const body = {
		magic_tag: "ss_app_log",
		header: {
			display_name: "TikTok",
			update_version_code: 2019092901,
			manifest_version_code: 2019092901,
			aid: 1233,
			channel: "googleplay",
			appkey: "5559e28267e58eb4c1000012",
			package: "com.zhiliaoapp.musically",
			app_version: "13.2.11",
			version_code: 130211,
			sdk_version: "2.5.5.8",
			os: "Android",
			os_version: "7.1.2",
			os_api: "25",
			device_model: this.device.default.device_type,
			device_brand: this.device.default.device_brand,
			cpu_abi: "arm64-v8a",
			release_build: "eaeeb2f_20190929",
			density_dpi: this.device.default.dpi,
			display_density: "mdpi",
			resolution: this.device.default.resolution,
			language: "en",
			mc: this.device.mac,
			timezone: 1,
			access: "wifi",
			not_request_sender: 0,
			carrier: this.device.carrier[2],
			mcc_mnc: this.device.default.mcc_mnc,
			google_aid: this.device.google_aid,
			openudid: this.device.default.openudid,
			clientudid: this.device.clientudid,
			sim_serial_number: [],
			tz_name: "America\\/New_York",
			tz_offset: 0,
			sim_region: "us"
		},
		_gen_time: Date.now()
	}
	const headers = {
		host: "log2.musical.ly",
		"sdk-version": "1",
		"content-type": "application/json; charset=utf-8",
		"user-agent": this.device["user-agent"]
	}
      /*
	const signature = await signRequest(requestUrl.protocol + requestUrl.host + requestUrl.pathname, requestUrl.search, body, headers);
	if (!signature) {
		throw new Error("No signature")
	} else {
		headers["X-Gorgon"] = signature["X-Gorgon"]
		headers["X-Khronos"] = signature["X-Khronos"]
      */
		const options = {
			method: "POST",
			url: requestUrl.href,
			headers: headers,
			body: body,
			json: true
		};
         console.log(requestUrl.href);
         console.log(headers);
         console.log(JSON.stringify(body));
}

async function signRequest(requestUrl, query, body, headers) {
	return new Promise(resolve => {
		const options = {
			method: "POST",
			url: "http://192.168.0.1:4000",
			headers: { "content-type": "application/json" },
			body: { url: requestUrl, query: query, body: body, headers: headers },
			json: true
		}
		request(options, function(error, response, body) {
			if (error) throw new Error(error)
			resolve(body)
		})
	})
}

function buildUrl(protocol, hostname, path, query) {
	if (hostname === "log2.musical.ly" || hostname === "api2.musical.ly") {
		query["_rticket"] = Date.now() * 1000
		query["ts"] = Date.now()
	}
	return url.parse(
		url.format({
			protocol: protocol,
			hostname: hostname,
			pathname: path,
			query: query
		})
	)
}

function getRandFromFile(path) {
	const file = fs.readFileSync(path)
	const data = JSON.parse(file)
	return data[Math.floor(Math.random() * data.length)]
}

function randomString(length, chars) {
	let mask = ""
	if (chars.indexOf("a") > -1) mask += "abcdefghijklmnopqrstuvwxyz"
	if (chars.indexOf("A") > -1) mask += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if (chars.indexOf("#") > -1) mask += "0123456789"
	let result = ""
	for (var i = length; i > 0; --i) result += mask[Math.floor(Math.random() * mask.length)]
	return result
}

function timeout(ms) {
	return new Promise(resolve => setTimeout(resolve, ms))
}

module.exports = TikTok
