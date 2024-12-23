9 llaves
```js
submit:function(){var t=this;if("undefined"!=typeof cms&&cms)return!1;if(!this.ajax)return this.$form.submit(),this;if(!this.isSubmitting){this.isSubmitting=!0;var e=$('input[type="file"]:not([disabled])',this.$form);e.each((function(t,e){e.files.length>0||$(e).prop("disabled",!0)}));var n=new FormData(this.$form[0]);"undefined"!=typeof webcard&&"culture"in webcard&&webcard.culture&&n.append("culture",webcard.culture),e.prop("disabled",!1),n.append("id",this.getId()),n.append("tac",Math.floor(((new Date).getTime()-this.time.getTime())/1e3));var r=this.$form.attr("action")||("undefined"!=typeof webcard&&webcard.apiHost?"https://"+webcard.apiHost:"")+"/form_container/submit";return $.ajax({url:r,crossDomain:!0,xhrFields:{withCredentials:!0},data:n,processData:!1,contentType:!1,type:"POST",success:function(e){t.isSubmitting=!1,t.$element.find(".ed-element").each((function(){this.element.destroy()})),t.$element.html(e),t.getViewport().scrollTo(t.$element,"center",300),$(e).is(".wv-success")?t.getViewport().notify("form.success",[t,t.$form]):t.getViewport().notify("form.error",[t,t.$form]),t.wakeup();var n=t.$element.find(".ed-element"),r=[];n.each((function(t,e){var n=i.Z.getTypeClassName($(e).attr("class"));n in window&&r.push(new window[n](e))})),animations.initSubsequent(r)}}).fail((function(e){t.isSubmitting=!1,t.$form.find(".wv-message").remove();var n=$('<div class="wv-message wv-failure">'.concat((null==e?void 0:e.responseText)||"Error while sending the form","</div>"));t.$form.prepend(n).append(n.clone())})),this}}
```

```js
submit: function() {
    var t = this;
    if ("undefined" != typeof cms && cms) return !1;
    if (!this.ajax) return this.$form.submit(), this;
    if (!this.isSubmitting) {
        this.isSubmitting = !0;
        var e = $('input[type="file"]:not([disabled])', this.$form);
        e.each((function(t, e) {
            e.files.length > 0 || $(e).prop("disabled", !0)
        }));
        var n = new FormData(this.$form[0]);
        "undefined" != typeof webcard && "culture" in webcard && webcard.culture && n.append("culture", webcard.culture), e.prop("disabled", !1), n.append("id", this.getId()), n.append("tac", Math.floor(((new Date).getTime() - this.time.getTime()) / 1e3));
        var r = this.$form.attr("action") || ("undefined" != typeof webcard && webcard.apiHost ? "https://" + webcard.apiHost : "") + "/form_container/submit";
        return $.ajax({
            url: r,
            crossDomain: !0,
            xhrFields: {
                withCredentials: !0
            },
            data: n,
            processData: !1,
            contentType: !1,
            type: "POST",
            success: function(e) {
                t.isSubmitting = !1, t.$element.find(".ed-element").each((function() {
                    this.element.destroy()
                })), t.$element.html(e), t.getViewport().scrollTo(t.$element, "center", 300), $(e).is(".wv-success") ? t.getViewport().notify("form.success", [t, t.$form]) : t.getViewport().notify("form.error", [t, t.$form]), t.wakeup();
                var n = t.$element.find(".ed-element"),
                    r = [];
                n.each((function(t, e) {
                    var n = i.Z.getTypeClassName($(e).attr("class"));
                    n in window && r.push(new window[n](e))
                })), animations.initSubsequent(r)
            }
        }).fail((function(e) {
            t.isSubmitting = !1, t.$form.find(".wv-message").remove();
            var n = $('<div class="wv-message wv-failure">'.concat((null == e ? void 0 : e.responseText) || "Error while sending the form", "</div>"));
            t.$form.prepend(n).append(n.clone())
        })), this
    }
}
```
