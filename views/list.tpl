<span id="addRecPanel" ng-show="!rCont.recFormShow">
    <span ng-repeat="reci in recipes">
        <span>{{"{{reci.name}}"}}</span>
        <span ng-repeat="ingr in reci.ingredients">
            {{"{{ingr.quantity}}"}} {{"{{ingr.unit}}"}} {{"{{ingr.ingredient.name}}"}}
        </span>
        <span>{{"{{reci.instructions.details}}"}}</span>
    </span>
</span>
