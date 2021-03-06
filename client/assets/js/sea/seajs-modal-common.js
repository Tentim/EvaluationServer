/*! 作者:阿伟 */



/*! git:https://github.com/aweiu/JsLibs.git */
/*! 推荐sealoader模块加载器:https://www.npmjs.com/package/sealoader */
/*! 最后修改于 2016-04-18 11:35:42 */
define(
    function (require, exports, module) {
        var assetsUrl = module.uri;

        assetsUrl = assetsUrl.substring(0, assetsUrl.lastIndexOf("/js/")) + "/";

        var waiting_bg, config, timer, timer_out, wrap_content, waiting_wrap, hideOnClickOut;

        exports["int"] = function (a) { config = a };

        var fixModal = new function () {

            var a;

            this.setOffset = function (b) { a = b, this.main() };

            this.main = function () {
                var b = .95 * document.documentElement.clientHeight;
                if (wrap_content.style.maxHeight = b + a + "px", config.hasBac === !1 && config.canClickOut) {
                    var c = waiting_bg.getBoundingClientRect();
                    waiting_bg.style.top = (document.documentElement.clientHeight - (c.bottom - c.top)) / 2 + "px", waiting_bg.style.left = (document.documentElement.clientWidth - (c.right - c.left)) / 2 + "px"
                }
            }
        };

        scrollCtrl = new function () {
            var a;

            this.enable = function () {
                a && (document.getElementsByTagName("html")[0].style.overflow = a)
            };

            this.dis = function () {
                if (config.hasBac !== !1 || !config.canClickOut) {
                    var b = document.getElementsByTagName("html")[0];
                    "hidden" != b.style.overflow && (a = getComputedStyle(b).overflow, b.style.overflow = "hidden");
                }
            }
        };

        HideOnClickOut = function () {
            var a = !0, b = function () { a ? a = !1 : exports.hide() }, c = function () { a = !0 }; this.enable = function () { setTimeout(function () { waiting_wrap.addEventListener("click", c), document.addEventListener("click", b) }, 0) }, this.dis = function () { waiting_wrap.removeEventListener("click", c), document.removeEventListener("click", b) }
        };

        exports.hide = function () {
            waiting_bg && (waiting_bg.style.display = "none", timer_out && clearTimeout(timer_out), timer && clearInterval(timer), scrollCtrl.enable(), hideOnClickOut && hideOnClickOut.dis(), config.afterHide && config.afterHide())
        };

        var autoHide = function () { var a = 1300; waiting_bg.style.opacity = 1, timer_out = setTimeout(function () { timer = setInterval(function () { return waiting_bg.style.opacity <= 0 ? (clearInterval(timer), void exports.hide()) : void (waiting_bg.style.opacity -= 1 / (a / 16)) }, 16) }, 1e3) };

        exports.show = function () {
            if (exports.hide(), null != waiting_bg) return waiting_wrap.style.behavior = "url(" + assetsUrl + "htc/PIE.htc)", waiting_bg.style.display = "table", document.body.appendChild(waiting_bg), scrollCtrl.dis(), config.afterShow && config.afterShow(wrap_content), config.autoHide && autoHide(), void (hideOnClickOut && hideOnClickOut.enable());
            if (!config.content) return void console.log("please int me..");
            waiting_bg = document.createElement("table"), waiting_bg.style.cssText = "position:fixed;z-index:" + (config.zIndex || 99999999999) + ";", config.hasBac !== !1 ? waiting_bg.style.cssText += ";top:0;left:0;height:100%;width:100%;background-color: rgba(0, 0, 0, .5);filter:progid:DXImageTransform.Microsoft.gradient(startColorstr=#7f000000,endColorstr=#7f000000);" : config.canClickOut || (waiting_bg.style.cssText += ";top:0;left:0;height:100%;width:100%;"), document.body.appendChild(waiting_bg); var a = document.createElement("td"); if (waiting_wrap = document.createElement("div"), a.style.cssText = "color:white;text-align:center;vertical-align: middle;", waiting_wrap.style.cssText = "display:inline-block;max-width:95%;", wrap_content = document.createElement("div"), config.hideOnClickOut && (hideOnClickOut = new HideOnClickOut), waiting_wrap.appendChild(wrap_content), wrap_content.style.cssText = "overflow:hidden;overflow-y:auto;", fixModal.setOffset(0), config.overflow && (wrap_content.style.overflow = config.overflow), null == config.hasWrap || config.hasWrap) { if (waiting_wrap.style.cssText += ";background-color: white;padding: 50px 0;border-radius: 10px;position: relative;overflow: hidden;", wrap_content.style.cssText += ";padding:2px 90px;", fixModal.setOffset(-100), config.padding && (waiting_wrap.style.padding = config.padding, wrap_content.style.paddingLeft = waiting_wrap.style.paddingLeft, waiting_wrap.style.paddingLeft = "0", wrap_content.style.paddingRight = waiting_wrap.style.paddingRight, waiting_wrap.style.paddingRight = "0"), config.hasShadow && (waiting_wrap.style.cssText += ";box-shadow:0px 0px 10px #666;"), config.hasTitleBar) { var b = waiting_wrap.style.paddingTop.replace("px", "") / 1, c = waiting_wrap.style.paddingBottom.replace("px", "") / 1; waiting_wrap.style.paddingTop = b + 45 + "px", fixModal.setOffset(-b - c - 45); var d, e = document.createElement("div"); if (e.style.cssText = "position:absolute;left: 0;top:0;background-color:#f7f7f7;height:45px;width:100%;color:#999;line-height:43px;border-radius: 10px 10px 0px 0px;", config.title) { var f = document.createElement("div"); f.style.cssText = "float:left;margin-left:15px;", f.innerHTML = config.title, e.appendChild(f) } waiting_wrap.appendChild(e) } config.hasClose && (d = document.createElement("div"), d.style.cssText = "cursor: pointer;font-size:30px;margin-right:15px;", e ? (d.style.cssText += ";float:right;", e.appendChild(d)) : (d.style.cssText += ";position:absolute;color:#dadada;", config.closePosition ? d.style.cssText += ";" + config.closePosition : d.style.cssText += ";right: 15px;top:5px"), d.innerHTML = "×", d.addEventListener("click", exports.hide), e || waiting_wrap.appendChild(d)) } config.textAlign && (wrap_content.style.textAlign = config.textAlign), a.appendChild(waiting_wrap), waiting_bg.appendChild(a); try { wrap_content.appendChild(config.content) } catch (g) { wrap_content.innerHTML += config.content } finally { scrollCtrl.dis(), config.afterShow && config.afterShow(wrap_content), config.autoHide && autoHide(), "undefined" != typeof myPlaceHolder && myPlaceHolder.main(wrap_content), waiting_wrap.style.behavior = "url(" + assetsUrl + "htc/PIE.htc)", fixModal.main(), window.addEventListener("resize", fixModal.main), hideOnClickOut && hideOnClickOut.enable() }
        }, exports.newInstance = function (fuc) { return function () { if (fuc += "", -1 == fuc.indexOf("\nexports=this;")) { fuc = ("0,(" + fuc + ")").replace("{", "{\nexports=this;"); var p = /([\s;=]+require\s*\([\s\S]*?\))/g; fuc = fuc.replace(p, "$1.newInstance()"), p = /( *\. *newInstance *\( *\)){2}/g, fuc = fuc.replace(p, ".newInstance()") } return new (eval(fuc))(require, exports, module) } }(arguments.callee)
   
    });