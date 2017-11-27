$(document).ready(function() {
  // Change page
  $('aside nav ul li a').click(function() {
    var current = $(this).attr('href').substr(1);
    $('aside nav ul li').removeClass('active');
    $('aside nav ul li').find('a[href$=' + current + ']').parent().addClass('active');
    $('main article:not([data-page="' + current + '"])').hide();
    $('main article[data-page="' + current + '"]').show();
  });
  // End change page

  // Change table in card1 row
  $('article[data-page="page3"] .card1 .ptop > a').click(function() {
    var currentT = $(this).attr('href').substr(1);
    console.log(currentT);
    $('article[data-page="page3"] .card1 .ptop a').removeClass('active');
    $('article[data-page="page3"] .card1 .ptop a[href$=' + currentT + ']').addClass('active');
    $('article[data-page="page3"] .card1 table:not([data-table="' + currentT + '"])').hide();
    $('article[data-page="page3"] .card1 table[data-table="' + currentT + '"]').show();
  });
  // End change table

  // Change table in card2 row
  $('article[data-page="page3"] .card2 .ptop > a').click(function() {
    var currentT = $(this).attr('href').substr(1);
    console.log(currentT);
    $('article[data-page="page3"] .card2 .ptop a').removeClass('active');
    $('article[data-page="page3"] .card2 .ptop a[href$=' + currentT + ']').addClass('active');
    $('article[data-page="page3"] .card2 table:not([data-table="' + currentT + '"])').hide();
    $('article[data-page="page3"] .card2 table[data-table="' + currentT + '"]').show();
  });
  // End change table
});
