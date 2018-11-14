<span id="listRecPanel" ng-show="!rCont.recFormShow && !rCont.delRecShow">
    <span ng-repeat="(ind, reci) in recipes">
        <span>{{"{{reci.name}}"}}</span>
        <span ng-repeat="ingr in reci.ingredients">
            {{"{{ingr.quantity}}"}} {{"{{ingr.unit}}"}} {{"{{ingr.ingredient.name}}"}}
        </span>
        <span>{{"{{reci.instructions.details}}"}}</span>
        <span ng-click="rCont.ShowDelConf(ind, reci)">X</span>
    </span>
</span>
