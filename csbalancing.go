package csbalancing

import (
	"sort"
)

// Entity has the ID and score of each client and CS
type Entity struct {
	ID    int
	Score int
}

// totalClientsPerEntity has the Entity data plus the total client of each CS
type totalClientsPerEntity struct {
	Entity
	totalClients int
}

type ByScore []totalClientsPerEntity

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// CustomerSuccessBalancing returns the customer success with more clients
func CustomerSuccessBalancing(customerSuccess []Entity, customers []Entity, customerSuccessAway []int) int {
	// Remove CSs on vacation
	avaliableCustomersSuccessess := getCustomerSuccessWorking(customerSuccess, customerSuccessAway)

	// Order CSs by score
	sort.Sort(ByScore(avaliableCustomersSuccessess))

	// set CS for each client
	customersByCustomerSuccess := setCustomerSuccessForEachClient(avaliableCustomersSuccessess, customers)

	return getTheCustommerSuccessWithMoreClients(customersByCustomerSuccess)
}

// getCustomerSuccessWorking removes the customer success that are away
func getCustomerSuccessWorking(customersSuccesses []Entity, customerSuccessAway []int) []totalClientsPerEntity {
	var avaliableCustomersSuccessess []totalClientsPerEntity
	for _, customerSuccess := range customersSuccesses {
		onVacation := false
		for _, csAway := range customerSuccessAway {
			if csAway == customerSuccess.ID {
				onVacation = true
			}
		}
		if !onVacation {
			// Creates a new struct with the total clients of each CS
			avaliableCustomersSuccessess = append(avaliableCustomersSuccessess, totalClientsPerEntity{customerSuccess, 0})
		}
	}

	return avaliableCustomersSuccessess
}

// setCustomerSuccessForEachClient add the client to a specific customer success
func setCustomerSuccessForEachClient(customerSuccess []totalClientsPerEntity, customers []Entity) []totalClientsPerEntity {
	for x := 0; x < len(customerSuccess); x++ {
		for i := len(customers) - 1; i >= 0; i-- {
			if customers[i].Score <= customerSuccess[x].Score {
				// Add +1 to total clients of CS and remove de client from the list to prevent of being added to 2 CSs
				customerSuccess[x].totalClients = customerSuccess[x].totalClients + 1
				customers = append(customers[:i], customers[i+1:]...)
			}
		}
	}

	return customerSuccess
}

// getTheCustommerSuccessWithMoreClients get the CS with more clients. Returns 0 with there is more than 1
func getTheCustommerSuccessWithMoreClients(customersByCustomerSuccess []totalClientsPerEntity) int {
	biggestCSID := 0
	totalClientsBiggerCS := 0
	for _, avaliableCustomerSuccess := range customersByCustomerSuccess {
		// If there is no client to this CS, jump to the next CS
		if avaliableCustomerSuccess.totalClients == 0 {
			continue
		}

		// If 2 CS have the same amount of clients, returns immediately
		if avaliableCustomerSuccess.totalClients == totalClientsBiggerCS {
			return 0
		}

		// Searches for the CS with more clients
		if avaliableCustomerSuccess.totalClients > totalClientsBiggerCS {
			totalClientsBiggerCS = avaliableCustomerSuccess.totalClients
			biggestCSID = avaliableCustomerSuccess.ID
		}
	}

	return biggestCSID
}