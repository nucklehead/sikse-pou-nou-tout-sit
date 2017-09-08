/**
 * Created by pearljam on 7/21/17.
 */

var spntApp = angular.module('spntApp', ['gfl.textAvatar']);

// Define the `PhoneListController` controller on the `phonecatApp` module
spntApp.controller('OptionController', function PhoneListController($scope) {
  $scope.options = [
    {
      title: 'Nexus S',
      description: 'Fast just got faster with Nexus S.'
    }, {
      title: 'Motorola XOOM™ with Wi-Fi',
      description: 'The Next, Next Generation tablet.'
    }, {
      title: 'MOTOROLA XOOM™',
      description: 'The Next, Next Generation tablet.'
    }, {
      title: 'Motorola XOOM™ with Wi-Fi',
      description: 'The Next, Next Generation tablet.'
    }, {
      title: 'MOTOROLA XOOM™',
      description: 'The Next, Next Generation tablet.'
    }, {
      title: 'Motorola XOOM™ with Wi-Fi',
      description: 'The Next, Next Generation tablet.'
    }, {
      title: 'MOTOROLA XOOM™',
      description: 'The Next, Next Generation tablet.'
    }, {
      title: 'Motorola XOOM™ with Wi-Fi',
      description: 'The Next, Next Generation tablet.'
    }, {
      title: 'MOTOROLA XOOM™',
      description: 'The Next, Next Generation tablet.'
    }, {
      title: 'Motorola XOOM™ with Wi-Fi',
      description: 'The Next, Next Generation tablet.'
    }, {
      title: 'MOTOROLA XOOM™',
      description: 'The Next, Next Generation tablet.'
    }
  ];
});
