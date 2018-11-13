(function(){
	var app = angular.module('recipUtil', []);
	app.controller('recCont', ['$scope', '$http', function($scope, $http){
		this.recForm = {ingreds: []};
		this.recFormShow = false;

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

		this.AddRecipe = function(){
			let ingNames = [];
			let ingDets = [];
			for (let i = 0; i < this.recForm.ingreds.length; i++){
				ingNames.push({name: this.recForm.ingreds[i].name});
				ingDets.push({quantity: this.recForm.ingreds[i].quantity, unit: this.recForm.ingreds[i].unit});
			}
			let sendData = {
				recipe: {
					name: this.recForm.name
				},
				ingredients: ingNames,
				recipe_ingredients: ingDets,
				method: {
					details: this.recForm.instructions
				}
			};
			$http.post("/recipe/add", sendData).then(function(ret){
				if (ret.data.success){

				}
			});
			this.recForm = {ingreds: []};
			this.recFormShow = false;
		};

		this.AddRecIng = function(){
			this.recForm.ingreds.push(this.ingForm);
			this.ingForm = {};
		};

		//Navigate
		this.ShowRecipeForm = function(rec_id = null){
			if (rec_id != null){

			}
			this.recFormShow = true;
		};
    }]);
})();
