interface directive {
    update(val:String | Element):
        void;
}
interface avalonObject {

}
interface avalonLibraryObj{
    $init?: Function;
    $ready?: Function;
    $dispose?: Function;
}
interface avalonComponentObj{
    $replace: Boolean;// 真值时表示替换其容器
    $slot: String; //默认插入点的名字
    $template: String; //组件的模板
    $extend: String; //指定要继承的组件名
    $container: Element ;//插入元素的位置,//比如dialog就不一定在使用它的位置插入,通常放在body中
    $construct: Function ;//用于调整三个配置项的合并,默认是function:(a, b,c ){return avalon.mix(a, b,c)}
    $$template:Function ;//用于微调组件的模板,传入$template与当前元素
    $init: Function ;//刚开始渲染时调用的回调, 传入vm与当前元素
    $childReady: Function;// 当其子组件$ready完毕后, 会冒泡到当前组件触发的回调, 传入vm与当前元素
    $ready: Function ;//当组件的所有子组件都调用其$ready回调后才触发的回调, 传入vm与当前元素
    $dispose: Function;// 该组件被移出DOM树，并且元素不存在msRetain属性时的销毁回调, 传入vm与当前元素
}
interface  avalonEventHandle {
    (e?:Event):boolean|void;
}
interface avalonWatchHandle {
    (o:any, n:any):void;
}
interface avalonArray {
    ensure(target:any[], item:any):any[];
    removeAt(target:any[], index:number):boolean;
    remove(target:any[], item:any):boolean;
}
interface avalonViewModel extends avalonViewModelInit{
    //$events:Object,// 放置$watch回调与绑定对象
    $watch(pro:string,fn:avalonWatchHandle):avalonViewModel,//增强版$watch
    $unwatch(tpye?:string,fn?:avalonWatchHandle):avalonViewModel,//
    //$fire: 触发$watch回调
    $track?:string[],//一个数组,里面包含用户定义的所有键名
    //$active:boolean,false时防止依赖收集
    $model?:Object,//返回一个纯净的JS对象
    //$accessors:放置所有读写器的数据描述对象
    $up?:avalonViewModel,//返回其上级对象
    $pathname?:avalonViewModel,//返回此对象在上级对象的名字,注意,数组元素的$pathname为空字符串
}
interface avalonViewModelInit {
    $id:string,
    //$events:Object,// 放置$watch回调与绑定对象
    //$watch: 增强版$watch
    //$fire: 触发$watch回调
    //$track?:string[],//一个数组,里面包含用户定义的所有键名
    //$active:boolean,false时防止依赖收集
    //$model?:Object,//返回一个纯净的JS对象
    //$accessors:放置所有读写器的数据描述对象
    //$up?:avalonViewModel,//返回其上级对象
    //$pathname?:avalonViewModel,//返回此对象在上级对象的名字,注意,数组元素的$pathname为空字符串
    //=============================
    $computed?:Object,//计算属性
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
