import{L as N,r as d,a as S,o as s,c as i,e as t,b as a,w as n,t as p,f as I,ad as T,k as H,z as A,F as V,j as P,l as v,E as J,q as U,x as K,p as X,g as Z,h as L}from"./index-BZvC1LFB.js";import{E as ee,a as te}from"./el-tab-pane-BQs3YRhz.js";import{s as oe,g as ae}from"./game-CmvRbD0X.js";import{E as le}from"./el-dialog-xM_jmMew.js";import{E as ne,a as se}from"./el-form-item-Cd_XCOi4.js";import{E as re}from"./el-card-BDDDD4UF.js";import{o as ce,c as ie,a as ue}from"./k8s-CuxIk5YK.js";import{M as de}from"./index-D1gv0VDz.js";/* empty css                */import{_ as $}from"./_plugin-vue_export-helper-DlAUqK2U.js";import"./strings-B2eoNP8o.js";import"./refs-ibKUFCnL.js";import"./isUndefined-DCTLXrZ8.js";import"./castArray-sZ2D1bz-.js";import"./_baseClone-BeqKcFuR.js";import"./_initCloneObject-Biqzf-Ei.js";const me=g=>(X("data-v-382c675b"),g=g(),Z(),g),pe={id:"challenge-card-bg",class:"hover:animate-background rounded-xl bg-gradient-to-r from-green-300 via-blue-500 to-purple-600 p-0.5 shadow-xl transition hover:bg-[length:400%_400%] hover:shadow-sm hover:[animation-duration:_4s]"},ge={class:"bg-gray-100 text-gray-800 py-1 px-2 rounded-full font-semibold mb-2 shadow-sm uppercase"},fe={style:{color:"#8f8f8f"}},ve={style:{color:"coral"}},_e={style:{"font-size":"18px","font-weight":"bold",color:"#00e18b"}},he={key:0,class:"solve challenge-solved"},be=me(()=>t("span",{style:{color:"#00dd30"}},"✔",-1)),ye={key:1,class:"solve challenge-nosolved"},xe={style:{height:"64vh",overflow:"auto",padding:"0 12px"}},ke={class:"block text-gray-700 font-bold mb-2"},we={class:"block text-gray-700 font-bold mb-2"},qe={class:"block text-gray-700 font-bold"},Ce={key:0,class:"btn-list"},Fe={class:"text-gray-900 bg-gradient-to-r from-lime-200 via-lime-400 to-lime-500 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-lime-300 dark:focus:ring-lime-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center"},Ie={key:1,class:"text-white bg-gradient-to-r from-red-400 via-red-500 to-red-600 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-red-300 dark:focus:ring-red-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2"},Te={__name:"ChallengeCard",props:{data:Object,challengeType:String},emits:["update"],setup(g,{emit:k}){const w=g,e=N(()=>({...w.data})),q=d(null),_=d({queFlag:[{required:!0,message:"请输入flag",trigger:"blur"}]}),l=d({queType:e.value.queType,queName:e.value.queName,queDescribe:e.value.queDescribe,queMark:e.value.queMark,Files:e.value.files,imageUrl:e.value.imageUrl,queFlag:""}),C=()=>{r.value=!0,e.value.imageUrl&&D()},h=()=>{r.value=!1},r=d(!1),E=k,F=async()=>{q.value.validate(async c=>{if(c){console.log(e);const o={questionId:e.value.id,submitFlag:l.value.queFlag},m=await oe(o);m.code===0?(v({type:"success",message:"提交flag成功"}),h(),E("update")):v({type:"error",message:m.msg})}})},b=c=>{const o=`/api/${c.url}`,m=document.createElement("a");m.href=o,m.download=c.name,document.body.appendChild(m),m.click(),document.body.removeChild(m)},u=d({containerImage:e.value.imageUrl,innerPort:e.value.innerPort}),f=d({containerAddr:"",outPort:""}),B=async()=>{if(console.log(u.value),!u.value.containerImage){v.error("未注册靶机镜像地址，请联系管理员");return}if(!u.value.innerPort){v.error("未注册靶机内部运行端口，请联系管理员");return}const c={imageAddr:u.value.containerImage,innerPort:parseInt(u.value.innerPort)};console.log(c);const o=await ce(c);o.code===0&&(f.value=o.data,v.success(o.msg))},D=async()=>{const c={imageAddr:u.value.containerImage},o=await ie(c);o.code===0&&(f.value=o.data,v.success(o.msg))},j=async()=>{const c={imageAddr:u.value.containerImage},o=await ue(c);o.code===0&&(f.value=o.data,v.success(o.msg))},z=()=>{window.open("http://"+f.value.containerAddr+":"+f.value.outPort)};return(c,o)=>{const m=re,y=ne,M=J,G=S("Share"),O=U,Q=K,W=se,Y=le;return s(),i(V,null,[t("div",pe,[a(m,{"body-style":{padding:"0px"},class:"challenge-card",shadow:"always"},{default:n(()=>[t("div",{class:"card-text",onClick:C},[t("div",ge,p(e.value.queType),1),t("h4",fe,p(e.value.queName),1),t("p",ve,"分值: "+p(e.value.queMark),1),t("div",_e," 已解决人数: "+p(e.value.queSolvers),1),e.value.ifSolved?(s(),i("div",he,[I(" You Sovled: "),be])):T("",!0),e.value.ifSolved?T("",!0):(s(),i("div",ye," You Not Sovled "))])]),_:1})]),a(Y,{modelValue:r.value,"onUpdate:modelValue":o[1]||(o[1]=x=>r.value=x),title:e.value.queName,"show-close":!1},{footer:n(()=>[t("div",{class:"dialog-footer"},[t("button",{onClick:h,class:"text-white bg-gradient-to-r from-cyan-500 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-cyan-300 dark:focus:ring-cyan-800 font-medium rounded-lg text-sm px-5 py-1.5 text-center me-2 mb-2"}," 关闭 ")])]),default:n(()=>[t("div",xe,[a(W,{ref_key:"challengeForm",ref:q,rules:_.value,model:l.value,"label-width":"100px"},{default:n(()=>[a(y,{label:"题目类型:",prop:"queType"},{default:n(()=>[t("label",ke,p(l.value.queType),1)]),_:1}),a(y,{label:"题目描述: （markdown）",prop:"queDescribe",class:"text-left"},{default:n(()=>[a(H(de),{modelValue:l.value.queDescribe},null,8,["modelValue"])]),_:1}),a(y,{label:"题目分值:",prop:"queMark"},{default:n(()=>[t("label",we,p(l.value.queMark),1)]),_:1}),l.value.Files&&l.value.Files.length?(s(),A(y,{key:0,label:"附件:",prop:"Files"},{default:n(()=>[(s(!0),i(V,null,P(l.value.Files,(x,R)=>(s(),i("div",{key:R,class:"flex justify-between items-center"},[t("label",qe,p(x.name),1),a(M,{onClick:()=>b(x),class:"bg-green-500 hover:bg-green-700 text-white font-bold py-1 px-4 mx-2 rounded"},{default:n(()=>[I("下载")]),_:2},1032,["onClick"])]))),128))]),_:1})):T("",!0),l.value.imageUrl?(s(),A(y,{key:1,label:"靶机:",prop:"queMark"},{default:n(()=>[t("div",null,[f.value.containerAddr?(s(),i("div",Ce,[t("label",Fe,p(f.value.containerAddr+":"+f.value.outPort),1),t("button",{onClick:z,type:"button",class:"text-gray-900 bg-gradient-to-r from-teal-200 to-lime-200 hover:bg-gradient-to-l hover:from-teal-200 hover:to-lime-200 focus:ring-4 focus:outline-none focus:ring-lime-200 dark:focus:ring-teal-700 font-medium rounded-lg text-sm px-5 py-2.5 text-center"},[I(" Open "),a(O,{class:"#dc2626"},{default:n(()=>[a(G)]),_:1})])])):(s(),i("label",Ie," 靶机未开启 "))]),t("div",{class:"btn-list mt-1"},[t("a",{class:"group inline-block rounded-full bg-gradient-to-r from-pink-500 via-red-500 to-yellow-500 p-[2px] hover:text-white focus:outline-none focus:ring active:text-opacity-75",href:"#"},[t("span",{onClick:B,class:"block rounded-full bg-white px-4 py-1 text-sm font-medium group-hover:bg-transparent"}," 开启靶机 ")]),t("a",{class:"group inline-block rounded-full bg-gradient-to-r from-pink-500 via-red-500 to-yellow-500 p-[2px] hover:text-white focus:outline-none focus:ring active:text-opacity-75",href:"#"},[t("span",{onClick:j,class:"block rounded-full bg-white px-4 py-1 text-sm font-medium group-hover:bg-transparent"}," 删除靶机 ")]),t("a",{class:"group inline-block rounded-full bg-gradient-to-r from-pink-500 via-red-500 to-yellow-500 p-[2px] hover:text-white focus:outline-none focus:ring active:text-opacity-75",href:"#"},[t("span",{onClick:D,class:"block rounded-full bg-white px-4 py-1 text-sm font-medium group-hover:bg-transparent"}," 查询靶机状态 ")])])]),_:1})):T("",!0),a(y,{label:"解题flag:",prop:"queFlag"},{default:n(()=>[a(Q,{modelValue:l.value.queFlag,"onUpdate:modelValue":o[0]||(o[0]=x=>l.value.queFlag=x)},null,8,["modelValue"]),a(M,{onClick:F,class:"mt-2 text-white bg-gradient-to-br from-pink-500 to-orange-400 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-pink-200 dark:focus:ring-pink-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2"},{default:n(()=>[I("提 交")]),_:1})]),_:1})]),_:1},8,["rules","model"])])]),_:1},8,["modelValue","title"])],64)}}},Ee=$(Te,[["__scopeId","data-v-382c675b"]]),Se={id:"challenge",class:"flex flex-wrap"},Ve={__name:"index",props:{challengeType:String},setup(g){const k=d(""),w=d(!0),e=d([]),q=async()=>{const r=await ae();r.code===0?e.value=r.data.gameList:v({type:"error",message:r.msg}),w.value=!1},_=g,l=async()=>{q()};l(),L(()=>k.value,()=>{l()});const C=()=>{l()},h=N(()=>e.value.filter(r=>r.queType.toLowerCase()===_.challengeType.toLowerCase()||_.challengeType.toLowerCase()==="all"));return(r,E)=>(s(),i("div",Se,[(s(!0),i(V,null,P(h.value,(F,b)=>(s(),i("div",{class:"w-1/4 p-4",key:b},[a(Ee,{data:F,challengeType:g.challengeType,onUpdate:C},null,8,["data","challengeType"])]))),128))]))}},Ae={id:"challenges"},Pe={style:{height:"auto"},class:"container"},De={class:"title flex flex-col items-center"},Me={__name:"index",setup(g){const k=d(!0),w={All:"mobile-phone",Crypto:"lock",Web:"monitor",Misc:"connection",Pwn:"cpu",Reverse:"refresh"},e=d("All");return L(()=>e.value,()=>{}),(q,_)=>{const l=S("Aim"),C=U,h=S("dv-decoration9"),r=S("dv-decoration3"),E=ee,F=te;return s(),i("div",Ae,[t("div",Pe,[t("div",De,[a(h,{class:"text-green-300 font-semibold",style:{width:"150px",height:"40px"}},{default:n(()=>[t("h1",null,[a(C,null,{default:n(()=>[a(l)]),_:1}),I(" "+p(e.value),1)])]),_:1}),a(r,{style:{width:"100%",height:"20px"}})]),a(F,{"tab-position":"left",style:{height:"100%"},modelValue:e.value,"onUpdate:modelValue":_[0]||(_[0]=b=>e.value=b),type:"border-card"},{default:n(()=>[(s(),i(V,null,P(w,(b,u)=>a(E,{name:u,key:u},{label:n(()=>[t("span",null,p(u),1)]),_:2},1032,["name"])),64)),k.value?(s(),A(Ve,{key:0,challengeType:e.value,style:{"min-height":"400px"}},null,8,["challengeType"])):T("",!0)]),_:1},8,["modelValue"])])])}}},Xe=$(Me,[["__scopeId","data-v-9f589fb9"]]);export{Xe as default};
