export namespace main {
	
	export class CategorySubItem {
	    id: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new CategorySubItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}
	export class CategoryItem {
	    id: number;
	    name: string;
	    list: CategorySubItem[];
	
	    static createFrom(source: any = {}) {
	        return new CategoryItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.list = this.convertValues(source["list"], CategorySubItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class LiveCategory {
	    code: number;
	    message: string;
	    data: CategoryItem[];
	
	    static createFrom(source: any = {}) {
	        return new LiveCategory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], CategoryItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RoomStatus {
	    code: number;
	    message: string;
	    status: number;
	    area_id: number;
	    parent_area_id: number;
	    title: string;
	
	    static createFrom(source: any = {}) {
	        return new RoomStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.status = source["status"];
	        this.area_id = source["area_id"];
	        this.parent_area_id = source["parent_area_id"];
	        this.title = source["title"];
	    }
	}
	export class User {
	    room_id: number;
	    name: string;
	    avatar: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.room_id = source["room_id"];
	        this.name = source["name"];
	        this.avatar = source["avatar"];
	    }
	}

}

export namespace model {
	
	export class Data {
	    url: string;
	    qrcode_key: string;
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.qrcode_key = source["qrcode_key"];
	    }
	}
	export class QRCode {
	    code?: number;
	    ttl?: number;
	    message?: string;
	    data: Data;
	
	    static createFrom(source: any = {}) {
	        return new QRCode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.ttl = source["ttl"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], Data);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Rtmp {
	    addr: string;
	    code: string;
	    new_link: string;
	    provider: string;
	
	    static createFrom(source: any = {}) {
	        return new Rtmp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.addr = source["addr"];
	        this.code = source["code"];
	        this.new_link = source["new_link"];
	        this.provider = source["provider"];
	    }
	}

}

