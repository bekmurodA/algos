#include <iostream>

int main(){
	std::cout<<"how much is your salary: ";
	int salary;
	std::cin>>salary;
	salary *=12; 
	std::cout<<"enter the yearly percent: ";
	int percent;
	std::cin>>percent;
	double const Percent = ((double)percent)/100 + 1;
	std::cout<<"enter the how much you want to achieve: ";
	int goal;
	std::cin>>goal;
	double amount = 0;	
	int years = 0 ;
	while(amount<=goal){
		++years;
		amount = amount*Percent+ salary;
		std::cout<<"AMOUNT IN YEAR" <<years<<": "<<amount<<"\n";
	}
	std::cout<<"it took "<<years<<"\n";
}

