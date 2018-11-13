(function(){
	var app = angular.module('recipUtil', []);
	app.controller('recCont', ['$scope', '$http', function($scope, $http){
		this.recForm = {ingreds: []};

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
		};

		this.AddRecIng = function(){
			this.recForm.ingreds.push(this.ingForm);
			this.ingForm = {};
		};
    }]);
})();
