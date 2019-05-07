var $context = null;

var simpleEditor = {
        /**
         * 初始化编辑器
         * @param $obj
         */
        init: function ($obj) {
            $context = document.createElement("div");
            $context.contentEditable = true;
            $context.className = $obj.className;
            $obj.style.display = "none";

            $obj.parentNode.insertBefore($context, $obj.nextSibling);

            $p = document.createElement("p");
            $p.style.color = "#B9B9B9";
            $p.innerText = "这一刻的想法...";

            $context.appendChild($p);
            $context.addEventListener("focus",function () {
                if($context.childNodes.length > 0){
                    for($i = 0;$i < $context.childNodes.length;$i++){
                        if($context.childNodes[$i] == $p){
                            $context.removeChild($p);
                            return false;
                        }
                    }
                }

            });
            $context.addEventListener("blur",function () {

                if ($context.innerHTML.length == 0){
                    $p.innerText = "这一刻的想法...";
                    $context.appendChild($p);
                }
            });

        },
        /**
         * 插入元素
         * @param obj 通常是一个node节点
         */
        insert: function (obj) {
            $context.appendChild(obj);
            $context.focus();
        },
        /**
         * 插入全部,一般用于编辑器初始化使用
         * @param obj 通常是html代码
         */
        insertAll: function (obj) {
            $context.innerHTML = obj;
            $context.focus();
        },
        /**
         * 取html代码内容
         * @returns {*}
         */
        getContextForHtml: function () {
            return $context.innerHTML;
        },
        /**
         * 取文本内容
         * @returns {UE.uNode.innerText|UM.uNode.innerText|string|*}
         */
        getContextForTxt: function () {
            return $context.innerText;
        },
        clear:function () {
            $context.innerHTML = "";
            $context.blur();
        }
}