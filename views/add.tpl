<span id="addRecPanel">
    <span>
        <form name="addRecipForm" id="addRecipForm" class="" novalidate>
            <span class=""><span>Name:</span><span class=""><input type="text" name="recName" id="recName" class="" ng-model="rCont.recForm.name" placeholder="Name" required/></span></span>
            <span class="" ng-show="rCont.recForm.ingreds.length > 0">
                <span ng-repeat="ingred in rCont.recForm.ingreds">
                    <span>{{"{{ingred.quantity}}"}} {{"{{ingred.unit}}"}} {{"{{ingred.name}}"}}</span>
                </span>
            </span>
            <span class=""><span>Instructions:</span><span class=""><textarea name="recInstr" id="recInstr" class="" ng-model="rCont.recForm.instructions" required></textarea></span></span>
            <input ng-disabled="!addRecipForm.$valid || rCont.recForm.ingreds.length == 0" ng-click="rCont.AddRecipe()" type="submit" />
        </form>
    </span>
    <span>
        <span>Ingredient:</span>
        <form name="addRecIngForm" id="" class="addRecIngForm" novalidate>
            <span>Quantity:</span><span class=""><input type="text" name="recQuan" id="recQuan" class="" ng-model="rCont.ingForm.quantity" placeholder="0" required/></span></span>
            <span>Unit:</span><span class=""><input type="text" name="recUnit" id="recUnit" class="" ng-model="rCont.ingForm.unit" placeholder="'Cups'" required/></span></span>
            <span>Name:</span><span class=""><input type="text" name="recName" id="recName" class="" ng-model="rCont.ingForm.name" placeholder="Name" required/></span></span>
            <input ng-disabled="!addRecIngForm.$valid" ng-click="rCont.AddRecIng()" type="submit" />
        </form>
    </span>
</span>
