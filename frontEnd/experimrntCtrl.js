app.controller('experimrntCtrl', function($scope, $location) {
    $scope.click = [];
    $scope.firstPlayer = false;
    $scope.mouseLocation = function(myE) {
        $scope.x = myE.clientX;
        $scope.y = myE.clientY;
    }
    $scope.mouseClick = function(i, j) {
        var loc = {x:i, y:j, firstPlayer:$scope.firstPlayer};
        $scope.click.push(loc)
        $scope.firstPlayer = !$scope.firstPlayer
        console.log($location)
    }
});