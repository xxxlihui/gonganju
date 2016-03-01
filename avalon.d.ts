interface directive {
    update(val:String | Element):
        void;
}
interface avalonObject {

}
interface  avalonEventHandle {
    (e?:Event):boolean|void;
}
interface avalonArray {
    ensure(target:any[], item:any):any[];
    removeAt(target:any[], index:number):boolean;
    remove(target:any[], item:any):boolean;
}
interface avalonViewModel {
    $id:string,
    //$events:Object,// 放置$watch回调与绑定对象
    //$watch: 增强版$watch
    //$fire: 触发$watch回调
    $track?:string[],//一个数组,里面包含用户定义的所有键名
    //$active:boolean,false时防止依赖收集
    $model?:Object,//返回一个纯净的JS对象
    //$accessors:放置所有读写器的数据描述对象
    $up?:avalonViewModel,//返回其上级对象
    $pathname?:avalonViewModel,//返回此对象在上级对象的名字,注意,数组元素的$pathname为空字符串
    //=============================
    $skipArray?:string[],//用于指定不可监听的属性,但VM生成是没有此属性的
}
interface avalonViewModelInit {
    $id:string,
    //$events:Object,// 放置$watch回调与绑定对象
    //$watch: 增强版$watch
    //$fire: 触发$watch回调
    $track?:string[],//一个数组,里面包含用户定义的所有键名
    //$active:boolean,false时防止依赖收集
    $model?:Object,//返回一个纯净的JS对象
    //$accessors:放置所有读写器的数据描述对象
    $up?:avalonViewModel,//返回其上级对象
    $pathname?:avalonViewModel,//返回此对象在上级对象的名字,注意,数组元素的$pathname为空字符串
    //=============================
    $skipArray?:string[],//用于指定不可监听的属性,但VM生成是没有此属性的
}
interface avalonStatic {
    directive(String, directive):directive;
    (el:Element):avalonObject;
    error(str:String, e?:Error):void;
    mix(a:any, b?:any|Boolean, ...c:any[]):any;
    log(...a:any[]):void;
    isFunction(s:any):boolean;
    noop():void;
    ready(fn:avalonEventHandle):void;
    onObject(a:string|string[], val?:any):Object;
    type(a:any):string;
    isWindow(a:Object):boolean;
    isPlainObject(a:Object):boolean;
    slice(a:Object|any[], start?:number, end?:number):any[];
    range(start:Number, end?:Number, step?:Number):Number[];
    bind(el:Element, type:String, fn:avalonEventHandle, phase?:boolean):Function;
    unbind(el:Element, type:string, fn:avalonEventHandle, phase?:boolean):void;
    each(obj:Object|any[], fn:(k:number|any, any)=>boolean|void):void;
    defind(vm:avalonViewModel):avalonViewModel;
    scan(el:Element, vm:avalonViewModel):void;
    css(node:Element|avalonObject, name:string):string|number;
    css(node:Element|avalonObject, name:string, val:string|number):void;
    nextTick(fn:Function):void;
    contains(a:Element, b:Element):boolean;
    parseHTML(html:string):DocumentFragment;
    innerHTML(node:Element, html:string):void;
    clearHTML(node:Element):void;
    Array:avalonArray;
}
declare var avalon:avalonStatic;
