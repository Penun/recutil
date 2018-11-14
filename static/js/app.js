(function(){
	var app = angular.module('recipUtil', []);
	app.controller('recCont', ['$scope', '$http', function($scope, $http){
		this.recForm = {ingreds: []};
		this.recFormShow = false;
		this.delRecShow = false;
		this.pendingDelIndex = null;
		this.pendingDelRec = null;

		angular.element(document).ready(function(){
			$http.get("/recipe/10").then(function(ret){
				if (ret.data.success){
					for (let i = 0; i < ret.data.recipes.length; i++){
						ret.data.recipes[i].ingredients = ret.data.ingredients[i];
						ret.data.recipes[i].instructions = ret.data.instructions[i];
					}
					$scope.recipes = ret.data.recipes;
				}
			});
		});


		// Add
		this.AddRecipe = function(){
			let ingNames = [];
			let ingDets = [];
			for (let i = 0; i < this.recForm.ingreds.length; i++){
				ingNames.push({name: this.recForm.ingreds[i].ingredient.name});
				ingDets.push({quantity: this.recForm.ingreds[i].quantity, unit: this.recForm.ingreds[i].unit});
			}
			let tmpDat = new Date();
			let sendData = {
				recipe: {
					name: this.recForm.name
				},
				ingredients: ingNames,
				recipe_ingredients: ingDets,
				method: {
					details: this.recForm.instructions
				},
				temp_id: tmpDat.getTime()
			};
			$http.post("/recipe/add", sendData).then(function(ret){
				if (ret.data.success){
					for (let i = 0; i < $scope.recipes.length; i++){
						if ($scope.recipes[i].temp_id == ret.data.temp_id){
							$scope.recipes[i].id = ret.data.id;
							delete $scope.recipes[i].temp_id;
							break;
						}
					}
				}
			});
			$scope.recipes.push({
				name: this.recForm.name,
				ingredients: this.recForm.ingreds,
				instructions: {
					details: this.recForm.instructions
				},
				temp_id: sendData.temp_id
			});
			this.recForm = {ingreds: []};
			this.recFormShow = false;
		};

		this.AddRecIng = function(){
			this.recForm.ingreds.push(this.ingForm);
			this.ingForm = {};
		};
		// End

		// Delete
		this.DeleteRecipe = function(){
			let sendData = {
				id: this.pendingDelRec.id,
				index: this.pendingDelIndex
			};
			$http.post("/recipe/delete", sendData).then(function(ret){
				if (ret.data.success){
					if ($scope.recipes[ret.data.index].id == ret.data.id){
						$scope.recipes.splice(ret.data.index, 1);
					} else {
						for (let i = 0; i < $scope.recipes.length; i++){
							if ($scope.recipes[i].id == ret.data.id){
								$scope.recipes.splice(i, 1);
								break;
							}
						}
					}
				}
			});
			this.CancelDelete();
		};

		this.CancelDelete = function(){
			this.pendingDelIndex = null;
			this.pendingDelRec = null;
			this.delRecShow = false;
		};
		// End

		//Navigate
		this.ShowRecipeForm = function(rec_id = null){
			if (rec_id != null){

			}
			this.recFormShow = true;
		};

		this.ShowDelConf = function(ind = null, reci = null){
			this.pendingDelIndex = ind;
			this.pendingDelRec = reci;
			if (ind != null && reci != null){
				this.delRecShow = true;
			}
		};
    }]);
})();
