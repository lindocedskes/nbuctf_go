import{_ as C,Q as T,Y as O,r as _,a3 as E,J as S,a2 as b,k as f,bk as L,B as m,aj as t,aP as $,ap as o}from"./index-BZvC1LFB.js";import{a as M,u as I}from"./el-popper-CKkpa1WL.js";const h=T({inheritAttrs:!1});function P(n,l,a,s,i,d){return O(n.$slots,"default")}var w=C(h,[["render",P],["__file","collection.vue"]]);const B=T({name:"ElCollectionItem",inheritAttrs:!1});function v(n,l,a,s,i,d){return O(n.$slots,"default")}var A=C(B,[["render",v],["__file","collection-item.vue"]]);const K="data-el-collection-item",Y=n=>{const l=`El${n}Collection`,a=`${l}Item`,s=Symbol(l),i=Symbol(a),d={...w,name:l,setup(){const u=_(null),c=new Map;E(s,{itemMap:c,getItems:()=>{const r=f(u);if(!r)return[];const e=Array.from(r.querySelectorAll(`[${K}]`));return[...c.values()].sort((N,g)=>e.indexOf(N.ref)-e.indexOf(g.ref))},collectionRef:u})}},y={...A,name:a,setup(u,{attrs:c}){const p=_(null),r=S(s,void 0);E(i,{collectionItemRef:p}),b(()=>{const e=f(p);e&&r.itemMap.set(e,{ref:e,...c})}),L(()=>{const e=f(p);r.itemMap.delete(e)})}};return{COLLECTION_INJECTION_KEY:s,COLLECTION_ITEM_INJECTION_KEY:i,ElCollection:d,ElCollectionItem:y}},k=m({trigger:M.trigger,effect:{...I.effect,default:"light"},type:{type:t(String)},placement:{type:t(String),default:"bottom"},popperOptions:{type:t(Object),default:()=>({})},id:String,size:{type:String,default:""},splitButton:Boolean,hideOnClick:{type:Boolean,default:!0},loop:{type:Boolean,default:!0},showTimeout:{type:Number,default:150},hideTimeout:{type:Number,default:150},tabindex:{type:t([Number,String]),default:0},maxHeight:{type:t([Number,String]),default:""},popperClass:{type:String,default:""},disabled:{type:Boolean,default:!1},role:{type:String,default:"menu"},buttonProps:{type:t(Object)},teleported:I.teleported}),D=m({command:{type:[Object,String,Number],default:()=>({})},disabled:Boolean,divided:Boolean,textValue:String,icon:{type:$}}),U=m({onKeydown:{type:t(Function)}}),x=[o.down,o.pageDown,o.home],j=[o.up,o.pageUp,o.end],V=[...x,...j],{ElCollection:q,ElCollectionItem:z,COLLECTION_INJECTION_KEY:G,COLLECTION_ITEM_INJECTION_KEY:H}=Y("Dropdown");export{H as C,q as E,V as F,j as L,D as a,K as b,Y as c,k as d,z as e,U as f,G as g};
