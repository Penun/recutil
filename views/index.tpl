{{template "includes/header.tpl"}}
<body ng-controller="recCont as rCont" ng-cloak>
    <span>
        <span>
            {{template "list.tpl"}}
            {{template "add.tpl"}}
            {{template "delete.tpl"}}
        </span>
        <span ng-click="rCont.ShowRecipeForm()" ng-show="!rCont.recFormShow">+</span>
    </span>
</body>
</html>
