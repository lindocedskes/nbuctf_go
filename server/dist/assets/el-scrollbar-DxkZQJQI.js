import{B as W,_ as D,Q as P,J as G,R as J,aR as ae,r as h,L as C,bk as se,b4 as N,bl as I,o as L,z as M,w as Q,i as le,e as O,W as B,k as g,Z as A,ay as oe,T as ne,an as Z,c as ee,b as F,F as re,aj as ie,G as R,aI as K,h as U,au as ce,n as Y,a3 as ue,a4 as ve,a2 as fe,c1 as de,Y as me,A as pe,ad as he,bI as be,a7 as ye,b7 as ge}from"./index-BZvC1LFB.js";const T=4,Se={vertical:{offset:"offsetHeight",scroll:"scrollTop",scrollSize:"scrollHeight",size:"height",key:"vertical",axis:"Y",client:"clientY",direction:"top"},horizontal:{offset:"offsetWidth",scroll:"scrollLeft",scrollSize:"scrollWidth",size:"width",key:"horizontal",axis:"X",client:"clientX",direction:"left"}},we=({move:s,size:t,bar:a})=>({[a.size]:t,transform:`translate${a.axis}(${s}%)`}),$=Symbol("scrollbarContextKey"),ze=W({vertical:Boolean,size:String,move:Number,ratio:{type:Number,required:!0},always:Boolean}),Ee="Thumb",_e=P({__name:"thumb",props:ze,setup(s){const t=s,a=G($),o=J("scrollbar");a||ae(Ee,"can not inject scrollbar context");const r=h(),c=h(),u=h({}),v=h(!1);let n=!1,b=!1,y=Z?document.onselectstart:null;const l=C(()=>Se[t.vertical?"vertical":"horizontal"]),i=C(()=>we({size:t.size,move:t.move,bar:l.value})),m=C(()=>r.value[l.value.offset]**2/a.wrapElement[l.value.scrollSize]/t.ratio/c.value[l.value.offset]),p=d=>{var w;if(d.stopPropagation(),d.ctrlKey||[1,2].includes(d.button))return;(w=window.getSelection())==null||w.removeAllRanges(),_(d);const H=d.currentTarget;H&&(u.value[l.value.axis]=H[l.value.offset]-(d[l.value.client]-H.getBoundingClientRect()[l.value.direction]))},E=d=>{if(!c.value||!r.value||!a.wrapElement)return;const w=Math.abs(d.target.getBoundingClientRect()[l.value.direction]-d[l.value.client]),H=c.value[l.value.offset]/2,x=(w-H)*100*m.value/r.value[l.value.offset];a.wrapElement[l.value.scroll]=x*a.wrapElement[l.value.scrollSize]/100},_=d=>{d.stopImmediatePropagation(),n=!0,document.addEventListener("mousemove",S),document.addEventListener("mouseup",f),y=document.onselectstart,document.onselectstart=()=>!1},S=d=>{if(!r.value||!c.value||n===!1)return;const w=u.value[l.value.axis];if(!w)return;const H=(r.value.getBoundingClientRect()[l.value.direction]-d[l.value.client])*-1,x=c.value[l.value.offset]-w,te=(H-x)*100*m.value/r.value[l.value.offset];a.wrapElement[l.value.scroll]=te*a.wrapElement[l.value.scrollSize]/100},f=()=>{n=!1,u.value[l.value.axis]=0,document.removeEventListener("mousemove",S),document.removeEventListener("mouseup",f),j(),b&&(v.value=!1)},e=()=>{b=!1,v.value=!!t.size},k=()=>{b=!0,v.value=n};se(()=>{j(),document.removeEventListener("mouseup",f)});const j=()=>{document.onselectstart!==y&&(document.onselectstart=y)};return N(I(a,"scrollbarElement"),"mousemove",e),N(I(a,"scrollbarElement"),"mouseleave",k),(d,w)=>(L(),M(ne,{name:g(o).b("fade"),persisted:""},{default:Q(()=>[le(O("div",{ref_key:"instance",ref:r,class:B([g(o).e("bar"),g(o).is(g(l).key)]),onMousedown:E},[O("div",{ref_key:"thumb",ref:c,class:B(g(o).e("thumb")),style:A(g(i)),onMousedown:p},null,38)],34),[[oe,d.always||v.value]])]),_:1},8,["name"]))}});var V=D(_e,[["__file","thumb.vue"]]);const He=W({always:{type:Boolean,default:!0},minSize:{type:Number,required:!0}}),Te=P({__name:"bar",props:He,setup(s,{expose:t}){const a=s,o=G($),r=h(0),c=h(0),u=h(""),v=h(""),n=h(1),b=h(1);return t({handleScroll:i=>{if(i){const m=i.offsetHeight-T,p=i.offsetWidth-T;c.value=i.scrollTop*100/m*n.value,r.value=i.scrollLeft*100/p*b.value}},update:()=>{const i=o==null?void 0:o.wrapElement;if(!i)return;const m=i.offsetHeight-T,p=i.offsetWidth-T,E=m**2/i.scrollHeight,_=p**2/i.scrollWidth,S=Math.max(E,a.minSize),f=Math.max(_,a.minSize);n.value=E/(m-E)/(S/(m-S)),b.value=_/(p-_)/(f/(p-f)),v.value=S+T<m?`${S}px`:"",u.value=f+T<p?`${f}px`:""}}),(i,m)=>(L(),ee(re,null,[F(V,{move:r.value,ratio:b.value,size:u.value,always:i.always},null,8,["move","ratio","size","always"]),F(V,{move:c.value,ratio:n.value,size:v.value,vertical:"",always:i.always},null,8,["move","ratio","size","always"])],64))}});var Ce=D(Te,[["__file","bar.vue"]]);const ke=W({height:{type:[String,Number],default:""},maxHeight:{type:[String,Number],default:""},native:{type:Boolean,default:!1},wrapStyle:{type:ie([String,Object,Array]),default:""},wrapClass:{type:[String,Array],default:""},viewClass:{type:[String,Array],default:""},viewStyle:{type:[String,Array,Object],default:""},noresize:Boolean,tag:{type:String,default:"div"},always:Boolean,minSize:{type:Number,default:20},id:String,role:String,ariaLabel:String,ariaOrientation:{type:String,values:["horizontal","vertical"]}}),Re={scroll:({scrollTop:s,scrollLeft:t})=>[s,t].every(R)},Le="ElScrollbar",Be=P({name:Le}),Pe=P({...Be,props:ke,emits:Re,setup(s,{expose:t,emit:a}){const o=s,r=J("scrollbar");let c,u;const v=h(),n=h(),b=h(),y=h(),l=C(()=>{const e={};return o.height&&(e.height=K(o.height)),o.maxHeight&&(e.maxHeight=K(o.maxHeight)),[o.wrapStyle,e]}),i=C(()=>[o.wrapClass,r.e("wrap"),{[r.em("wrap","hidden-default")]:!o.native}]),m=C(()=>[r.e("view"),o.viewClass]),p=()=>{var e;n.value&&((e=y.value)==null||e.handleScroll(n.value),a("scroll",{scrollTop:n.value.scrollTop,scrollLeft:n.value.scrollLeft}))};function E(e,k){be(e)?n.value.scrollTo(e):R(e)&&R(k)&&n.value.scrollTo(e,k)}const _=e=>{R(e)&&(n.value.scrollTop=e)},S=e=>{R(e)&&(n.value.scrollLeft=e)},f=()=>{var e;(e=y.value)==null||e.update()};return U(()=>o.noresize,e=>{e?(c==null||c(),u==null||u()):({stop:c}=ce(b,f),u=N("resize",f))},{immediate:!0}),U(()=>[o.maxHeight,o.height],()=>{o.native||Y(()=>{var e;f(),n.value&&((e=y.value)==null||e.handleScroll(n.value))})}),ue($,ve({scrollbarElement:v,wrapElement:n})),fe(()=>{o.native||Y(()=>{f()})}),de(()=>f()),t({wrapRef:n,update:f,scrollTo:E,setScrollTop:_,setScrollLeft:S,handleScroll:p}),(e,k)=>(L(),ee("div",{ref_key:"scrollbarRef",ref:v,class:B(g(r).b())},[O("div",{ref_key:"wrapRef",ref:n,class:B(g(i)),style:A(g(l)),onScroll:p},[(L(),M(pe(e.tag),{id:e.id,ref_key:"resizeRef",ref:b,class:B(g(m)),style:A(e.viewStyle),role:e.role,"aria-label":e.ariaLabel,"aria-orientation":e.ariaOrientation},{default:Q(()=>[me(e.$slots,"default")]),_:3},8,["id","class","style","role","aria-label","aria-orientation"]))],38),e.native?he("v-if",!0):(L(),M(Ce,{key:0,ref_key:"barRef",ref:y,always:e.always,"min-size":e.minSize},null,8,["always","min-size"]))],2))}});var xe=D(Pe,[["__file","scrollbar.vue"]]);const Me=ye(xe),z=new Map;let X;Z&&(document.addEventListener("mousedown",s=>X=s),document.addEventListener("mouseup",s=>{for(const t of z.values())for(const{documentHandler:a}of t)a(s,X)}));function q(s,t){let a=[];return Array.isArray(t.arg)?a=t.arg:ge(t.arg)&&a.push(t.arg),function(o,r){const c=t.instance.popperRef,u=o.target,v=r==null?void 0:r.target,n=!t||!t.instance,b=!u||!v,y=s.contains(u)||s.contains(v),l=s===u,i=a.length&&a.some(p=>p==null?void 0:p.contains(u))||a.length&&a.includes(v),m=c&&(c.contains(u)||c.contains(v));n||b||y||l||i||m||t.value(o,r)}}const Oe={beforeMount(s,t){z.has(s)||z.set(s,[]),z.get(s).push({documentHandler:q(s,t),bindingFn:t.value})},updated(s,t){z.has(s)||z.set(s,[]);const a=z.get(s),o=a.findIndex(c=>c.bindingFn===t.oldValue),r={documentHandler:q(s,t),bindingFn:t.value};o>=0?a.splice(o,1,r):a.push(r)},unmounted(s){z.delete(s)}};export{Oe as C,Me as E};
