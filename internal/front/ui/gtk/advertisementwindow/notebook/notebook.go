package notebook

func (n *Notebook) SetSensetive(b bool) {
	n.notebook.SetSensitive(b)
}

func (n *Notebook) BlockAllSignals() {
	n.BlockAdvertisementsTab.BlockSignals()
	n.LineAdvertisementsTab.BlockSignals()
	n.OrdersTab.BlockSignals()
	n.ClientsTab.BlockSignals()
	n.ExtrachargesTab.BlockSignals()
	n.TagsTab.BlockSignals()
}

func (n *Notebook) DisableAllModels() {
	n.BlockAdvertisementsTab.UnsetModel()
	n.LineAdvertisementsTab.UnsetModel()
	n.OrdersTab.UnsetModel()
	n.ClientsTab.UnsetModel()
	n.ExtrachargesTab.UnsetModel()
	n.TagsTab.UnsetModel()
}

func (n *Notebook) EnableAllModels() {
	n.BlockAdvertisementsTab.SetModel()
	n.LineAdvertisementsTab.SetModel()
	n.OrdersTab.SetModel()
	n.ClientsTab.SetModel()
	n.ExtrachargesTab.SetModel()
	n.TagsTab.SetModel()
}

func (n *Notebook) UnblockAllSignals() {
	n.BlockAdvertisementsTab.UnblockSignals()
	n.LineAdvertisementsTab.UnblockSignals()
	n.OrdersTab.UnblockSignals()
	n.ClientsTab.UnblockSignals()
	n.ExtrachargesTab.UnblockSignals()
	n.TagsTab.UnblockSignals()
}

func (n *Notebook) ClearAllListStores() {
	n.lists.ClearBlockAdvertisementList()
	n.lists.ClearLineAdvertisementList()
	n.lists.ClearOrderList()
	n.lists.ClearClientList()
	n.lists.ClearTagList()
	n.lists.ClearExtraChargeList()
	n.lists.ClearCostRateListStore()
}

func (n *Notebook) EnableAdvertisementFilters(enable bool) {
	n.BlockAdvertisementsTab.SetEnableFilters(enable)
	n.LineAdvertisementsTab.SetEnableFilters(enable)
}

func (n *Notebook) EnableSidebarsFilters(enable bool) {

	n.ExtrachargesTab.SetEnableFilters(enable)
	n.TagsTab.SetEnableFilters(enable)
	n.ClientsTab.SetEnableFilters(enable)
	n.OrdersTab.SetEnableFilters(enable)
}

func (n *Notebook) ResetSorts() {
	n.BlockAdvertisementsTab.ResetSort()
	n.LineAdvertisementsTab.ResetSort()
	n.ExtrachargesTab.ResetSort()
	n.TagsTab.ResetSort()
	n.ClientsTab.ResetSort()
	n.OrdersTab.ResetSort()
}

func (n *Notebook) AttachAll() {
	n.BlockAdvertisementsTab.ResetSort()
	n.LineAdvertisementsTab.ResetSort()
	n.ExtrachargesTab.ResetSort()
	n.TagsTab.ResetSort()
	n.ClientsTab.ResetSort()
	n.OrdersTab.ResetSort()
}

func (n *Notebook) RefilterBlock() {
	n.BlockAdvertisementsTab.Refilter()
}

func (n *Notebook) RefilterLine() {
	n.LineAdvertisementsTab.Refilter()
}

func (n *Notebook) LockClientsTab(l bool) {
	n.ClientsTab.SetSensitive(!l)
}

func (n *Notebook) LockOrdersTab(l bool) {
	n.OrdersTab.SetSensitive(!l)
}

func (n *Notebook) LockBlockAdvertisementsTab(l bool) {
	n.BlockAdvertisementsTab.SetSensitive(!l)
}

func (n *Notebook) LockLineAdvertisementsTab(l bool) {
	n.LineAdvertisementsTab.SetSensitive(!l)
}

func (n *Notebook) LockTagsTab(l bool) {
	n.TagsTab.SetSensitive(!l)
}

func (n *Notebook) LockExtraChargesTab(l bool) {
	n.ExtrachargesTab.SetSensitive(!l)
}
