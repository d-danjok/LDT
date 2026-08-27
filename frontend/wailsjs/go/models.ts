export namespace main {
	
	export class AssemblyInfo {
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new AssemblyInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}
	export class LCVersionInfo {
	    name: string;
	    lastDate: string;
	
	    static createFrom(source: any = {}) {
	        return new LCVersionInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.lastDate = source["lastDate"];
	    }
	}

}

