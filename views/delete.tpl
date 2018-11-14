<span id="delRecPanel" ng-show="rCont.delRecShow">
    <span>Are you sure you would like to delete {{"{{rCont.pendingDelRec.name}}"}}?</span>
    <span ng-click="rCont.DeleteRecipe()">Yes</span><span ng-click="rCont.CancelDelete()">No</span>
</span>
